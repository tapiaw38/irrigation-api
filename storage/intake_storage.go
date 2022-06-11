package storage

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models/intake"
)

type IntakeStorage struct {
	Data *Data
}

// CreateIntake creates a new Intake in the database.
func (is *IntakeStorage) CreateIntakes(ctx context.Context, intakes []intake.Intake) ([]intake.Intake, error) {

	q := `
	INSERT INTO intakes (intake_number, name, section, latitude, longitude, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	its := []intake.Intake{}

	for _, i := range intakes {

		var id int64

		row := is.Data.DB.QueryRowContext(
			ctx, q,
			i.IntakeNumber,
			i.Name,
			i.Section,
			i.Latitude,
			i.Longitude,
			time.Now(),
			time.Now(),
		)

		err := row.Scan(&id)

		if err != nil {
			log.Println(err)
			return its, err
		}

		i.ID = id

		its = append(its, i)
	}

	return its, nil
}

// GetIntakes gets an intake from the database.
func (is *IntakeStorage) GetIntakes(ctx context.Context) ([]intake.IntakeResponse, error) {

	q := `
	SELECT intakes.id, sections.id, sections.section_number,
			sections.name,
			intakes.intake_number, intakes.name, intakes.latitude, 
			intakes.longitude, intakes.created_at, intakes.updated_at
		FROM intakes
		LEFT JOIN sections ON intakes.section = sections.id;
	`

	rows, err := is.Data.DB.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	intakes := []intake.IntakeResponse{}

	for rows.Next() {
		iks, err := ScanRowIntakeResponse(rows)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		intakes = append(intakes, iks)
	}

	return intakes, nil
}

// UpdateIntake updates an intake in the database.
func (is *IntakeStorage) UpdateIntake(ctx context.Context, id string, intake intake.Intake) (intake.IntakeResponse, error) {

	q := `
	WITH updated AS (
		UPDATE intakes
		SET intake_number = $1, name = $2, section = $3, latitude = $4, 
			longitude = $5, updated_at = $6
		WHERE id = $7
		RETURNING id, intakes.intake_number, intakes.name,
			intakes.section, intakes.latitude, intakes.longitude,
			intakes.created_at, intakes.updated_at
	)
	SELECT updated.id, sections.id, sections.section_number,
			sections.name,
			updated.intake_number, updated.name, updated.latitude,
			updated.longitude, updated.created_at, updated.updated_at
	FROM updated
	LEFT JOIN sections ON updated.section = sections.id;
	`

	row := is.Data.DB.QueryRowContext(
		ctx, q,
		intake.IntakeNumber,
		intake.Name,
		intake.Section,
		intake.Latitude,
		intake.Longitude,
		time.Now(),
		id,
	)

	iks, err := ScanRowIntakeResponse(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	return iks, nil
}

// DeleteIntake deletes an intake from the database.
func (is *IntakeStorage) DeleteIntake(ctx context.Context, id string) (intake.Intake, error) {

	q := `
	DELETE FROM intakes
		WHERE id = $1
		RETURNING id, intake_number, name, section, latitude, longitude, created_at, updated_at;
	`

	row := is.Data.DB.QueryRowContext(
		ctx, q,
		id,
	)

	iks, err := ScanRowIntake(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	return iks, nil
}
