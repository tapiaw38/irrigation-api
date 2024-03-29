package database

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

// CreateIntake creates a new Intake in the database.
func (is *PostgresRepository) CreateIntakes(ctx context.Context, intakes []models.Intake) ([]models.Intake, error) {

	q := `
	INSERT INTO intakes (intake_number, name, section, latitude, longitude, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	its := []models.Intake{}

	for _, i := range intakes {

		var id int64

		row := is.db.QueryRowContext(
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
func (is *PostgresRepository) GetIntakes(ctx context.Context) ([]models.IntakeResponse, error) {

	q := `
	SELECT intakes.id, sections.id, sections.section_number,
			sections.name,
			intakes.intake_number, intakes.name, intakes.latitude, 
			intakes.longitude, intakes.created_at, intakes.updated_at
		FROM intakes
		LEFT JOIN sections ON intakes.section = sections.id
		ORDER BY intakes.intake_number ASC;
	`

	rows, err := is.db.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	intakes := []models.IntakeResponse{}

	for rows.Next() {

		iks, err := ScanRowIntakeResponse(rows)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		q = `
		SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
			producers.document_number, producers.birth_date, producers.phone_number,
			producers.address,
			productions.lote_number, productions.entry, 
			productions.name, productions.production_type, productions.area, 
			productions.cultivated_area, productions.latitude, productions.longitude, 
			productions.picture, productions.cadastral_registration, productions.district,
			intakes_productions.watering_order, productions.created_at, productions.updated_at
			FROM productions
			LEFT JOIN producers ON productions.producer = producers.id
			LEFT JOIN intakes_productions
            ON intakes_productions.production_id=productions.id
            WHERE intakes_productions.intake_id=$1
		`

		rows, err := is.db.QueryContext(
			ctx, q,
			iks.ID,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		for rows.Next() {
			pds, err := ScanRowProductionIntakeResponse(rows)

			if err != nil {
				log.Println(err)
				return nil, err
			}

			iks.Productions = append(iks.Productions, pds)
		}

		intakes = append(intakes, iks)
	}

	return intakes, nil
}

// GetIntakeByID gets an intake from the database.
func (is *PostgresRepository) GetIntakeByID(ctx context.Context, id string) (models.IntakeResponse, error) {
	q := `
	SELECT intakes.id, sections.id, sections.section_number,
			sections.name,
			intakes.intake_number, intakes.name, intakes.latitude, 
			intakes.longitude, intakes.created_at, intakes.updated_at
		FROM intakes
		LEFT JOIN sections ON intakes.section = sections.id
		WHERE intakes.id = $1;
	`

	row := is.db.QueryRowContext(ctx, q, id)

	iks, err := ScanRowIntakeResponse(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes_productions.watering_order, productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		WHERE intakes_productions.intake_id=$1
		ORDER BY intakes_productions.watering_order ASC;
	`

	rows, err := is.db.QueryContext(
		ctx, q,
		iks.ID,
	)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	for rows.Next() {
		pds, err := ScanRowProductionIntakeResponse(rows)

		if err != nil {
			log.Println(err)
			return iks, err
		}

		iks.Productions = append(iks.Productions, pds)
	}

	return iks, nil
}

// UpdateIntake updates an intake in the database.
func (is *PostgresRepository) UpdateIntake(ctx context.Context, id string, intake models.Intake) (models.IntakeResponse, error) {

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

	row := is.db.QueryRowContext(
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
func (is *PostgresRepository) DeleteIntake(ctx context.Context, id string) (models.Intake, error) {

	q := `
	DELETE FROM intakes
		WHERE id = $1
		RETURNING id, intake_number, name, section, latitude, longitude, created_at, updated_at;
	`

	row := is.db.QueryRowContext(
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

// CreateIntakeProduction creates a intake production many-to-many relationship in the database.
func (is *PostgresRepository) CreateIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error) {

	var iks models.IntakeResponse

	q := `
	INSERT INTO intakes_productions (intake_id, production_id, watering_order)
		VALUES ($1, $2, $3)
		RETURNING intake_id, production_id, watering_order;
	`

	row := is.db.QueryRowContext(
		ctx, q,
		intakeID,
		intakeProduction.ProductionID,
		intakeProduction.WateringOrder,
	)

	ip, err := ScanRowIntakeProduction(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT intakes.id, sections.id, sections.section_number,
			sections.name,
			intakes.intake_number, intakes.name, intakes.latitude, 
			intakes.longitude, intakes.created_at, intakes.updated_at
		FROM intakes
		LEFT JOIN sections ON intakes.section = sections.id
		WHERE intakes.id = $1;
	`

	row = is.db.QueryRowContext(
		ctx, q,
		ip.IntakeID,
	)

	iks, err = ScanRowIntakeResponse(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes_productions.watering_order, productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		WHERE intakes_productions.intake_id=$1
		ORDER BY intakes_productions.watering_order ASC;
	`

	rows, err := is.db.QueryContext(
		ctx, q,
		iks.ID,
	)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	for rows.Next() {
		pds, err := ScanRowProductionIntakeResponse(rows)

		if err != nil {
			log.Println(err)
			return iks, err
		}

		iks.Productions = append(iks.Productions, pds)
	}

	return iks, nil
}

// UpdateIntakeProduction updates an intake production many-to-many relationship in the database.
func (is *PostgresRepository) UpdateIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error) {

	var iks models.IntakeResponse

	q := `
	UPDATE intakes_productions
		SET watering_order = $1
		WHERE intake_id = $2 AND production_id = $3
		RETURNING intake_id, production_id, watering_order;
	`

	row := is.db.QueryRowContext(
		ctx, q,
		intakeProduction.WateringOrder,
		intakeID,
		intakeProduction.ProductionID,
	)

	ip, err := ScanRowIntakeProduction(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT intakes.id, sections.id, sections.section_number,
			sections.name,
			intakes.intake_number, intakes.name, intakes.latitude, 
			intakes.longitude, intakes.created_at, intakes.updated_at
		FROM intakes
		LEFT JOIN sections ON intakes.section = sections.id
		WHERE intakes.id = $1;
	`

	row = is.db.QueryRowContext(
		ctx, q,
		ip.IntakeID,
	)

	iks, err = ScanRowIntakeResponse(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes_productions.watering_order, productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		WHERE intakes_productions.intake_id=$1
		ORDER BY intakes_productions.watering_order ASC;
	`

	rows, err := is.db.QueryContext(
		ctx, q,
		iks.ID,
	)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	for rows.Next() {
		pds, err := ScanRowProductionIntakeResponse(rows)

		if err != nil {
			log.Println(err)
			return iks, err
		}

		iks.Productions = append(iks.Productions, pds)
	}

	return iks, nil
}

// DeleteIntakeProduction deletes a intake production many-to-many relationship in the database.
func (is *PostgresRepository) DeleteIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error) {

	var iks models.IntakeResponse

	q := `
	DELETE FROM intakes_productions
		WHERE intake_id = $1 AND production_id = $2
		RETURNING intake_id, production_id, watering_order;
	`

	row := is.db.QueryRowContext(
		ctx, q,
		intakeID,
		intakeProduction.ProductionID,
	)

	ip, err := ScanRowIntakeProduction(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT intakes.id, sections.id, sections.section_number,
			sections.name,
			intakes.intake_number, intakes.name, intakes.latitude, 
			intakes.longitude, intakes.created_at, intakes.updated_at
		FROM intakes
		LEFT JOIN sections ON intakes.section = sections.id
		WHERE intakes.id = $1;
	`

	row = is.db.QueryRowContext(
		ctx, q,
		ip.IntakeID,
	)

	iks, err = ScanRowIntakeResponse(row)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes_productions.watering_order, productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		WHERE intakes_productions.intake_id=$1
		ORDER BY intakes_productions.watering_order ASC;
	`

	rows, err := is.db.QueryContext(
		ctx, q,
		iks.ID,
	)

	if err != nil {
		log.Println(err)
		return iks, err
	}

	for rows.Next() {
		pds, err := ScanRowProductionIntakeResponse(rows)

		if err != nil {
			log.Println(err)
			return iks, err
		}

		iks.Productions = append(iks.Productions, pds)
	}

	return iks, nil
}
