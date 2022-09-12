package database

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

// CreateSections creates a new Section in the database.
func (ss *PostgresRepository) CreateSections(ctx context.Context, sections []models.Section) ([]models.Section, error) {

	q := `
	INSERT INTO sections (section_number, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	sct := []models.Section{}

	for _, s := range sections {

		var id int64

		row := ss.db.QueryRowContext(
			ctx, q,
			s.SectionNumber,
			s.Name,
			time.Now(),
			time.Now(),
		)

		err := row.Scan(&id)

		if err != nil {
			log.Println(err)
			return sct, err
		}

		s.ID = id

		sct = append(sct, s)
	}

	return sct, nil

}

// GetSections returns all Sections from the database.
func (ss *PostgresRepository) GetSections(ctx context.Context) ([]models.Section, error) {

	q := `
	SELECT id, section_number, name, created_at, updated_at
		FROM sections
		ORDER BY section_number ASC;
	`

	rows, err := ss.db.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	sct := []models.Section{}

	for rows.Next() {
		s, err := ScanRowSection(rows)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		sct = append(sct, s)
	}

	return sct, nil
}

// GetSectionByID returns a Section from the database by id.
func (ss *PostgresRepository) GetSectionByID(ctx context.Context, id string) (models.Section, error) {

	q := `
	SELECT id, section_number, name, created_at, updated_at
		FROM sections
		WHERE id = $1;
	`

	row := ss.db.QueryRowContext(ctx, q, id)

	s, err := ScanRowSection(row)

	if err != nil {
		log.Println(err)
		return models.Section{}, err
	}

	return s, nil
}

// UpdateSection updates a Section in the database.
func (ss *PostgresRepository) UpdateSection(ctx context.Context, id string, s models.Section) (models.Section, error) {

	q := `
	UPDATE sections
		SET section_number = $1, name = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, section_number, name, created_at, updated_at;
	`

	row := ss.db.QueryRowContext(
		ctx, q,
		s.SectionNumber,
		s.Name,
		time.Now(),
		id,
	)

	s, err := ScanRowSection(row)

	if err != nil {
		log.Println(err)
		return models.Section{}, err
	}

	return s, nil
}

// DeleteSection deletes a Section from the database.
func (ss *PostgresRepository) DeleteSection(ctx context.Context, id string) (models.Section, error) {

	q := `
	DELETE FROM sections
		WHERE id = $1
		RETURNING id, section_number, name, created_at, updated_at;
	`

	row := ss.db.QueryRowContext(ctx, q, id)

	s, err := ScanRowSection(row)

	if err != nil {
		log.Println(err)
		return models.Section{}, err
	}

	return s, nil
}
