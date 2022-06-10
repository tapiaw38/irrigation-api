package storage

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models/section"
)

type SectionStorage struct {
	Data *Data
}

// CreateSections creates a new Section in the database.
func (ss *SectionStorage) CreateSections(ctx context.Context, sections []section.Section) ([]section.Section, error) {

	q := `
	INSERT INTO sections (section_number, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	sct := []section.Section{}

	for _, s := range sections {

		var id int64

		row := ss.Data.DB.QueryRowContext(
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
func (ss *SectionStorage) GetSections(ctx context.Context) ([]section.Section, error) {

	q := `
	SELECT id, section_number, name, created_at, updated_at
		FROM sections
		ORDER BY section_number ASC;
	`

	rows, err := ss.Data.DB.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	sct := []section.Section{}

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
func (ss *SectionStorage) GetSectionByID(ctx context.Context, id string) (section.Section, error) {

	q := `
	SELECT id, section_number, name, created_at, updated_at
		FROM sections
		WHERE id = $1;
	`

	row := ss.Data.DB.QueryRowContext(ctx, q, id)

	s, err := ScanRowSection(row)

	if err != nil {
		log.Println(err)
		return section.Section{}, err
	}

	return s, nil
}

// UpdateSection updates a Section in the database.
func (ss *SectionStorage) UpdateSection(ctx context.Context, id string, s section.Section) (section.Section, error) {

	q := `
	UPDATE sections
		SET section_number = $1, name = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, section_number, name, created_at, updated_at;
	`

	row := ss.Data.DB.QueryRowContext(
		ctx, q,
		s.SectionNumber,
		s.Name,
		time.Now(),
		id,
	)

	s, err := ScanRowSection(row)

	if err != nil {
		log.Println(err)
		return section.Section{}, err
	}

	return s, nil
}

// DeleteSection deletes a Section from the database.
func (ss *SectionStorage) DeleteSection(ctx context.Context, id string) (section.Section, error) {

	q := `
	DELETE FROM sections
		WHERE id = $1
		RETURNING id, section_number, name, created_at, updated_at;
	`

	row := ss.Data.DB.QueryRowContext(ctx, q, id)

	s, err := ScanRowSection(row)

	if err != nil {
		log.Println(err)
		return section.Section{}, err
	}

	return s, nil
}
