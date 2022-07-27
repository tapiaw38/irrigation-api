package storage

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models/producer"
)

type ProducerStorage struct {
	Data *Data
}

// CreateProducer creates a new producer in the database
func (ps *ProducerStorage) CreateProducers(ctx context.Context, producers []producer.Producer) ([]producer.Producer, error) {

	q := `
	INSERT INTO producers (first_name, last_name, document_number, birth_date, phone_number, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	pds := []producer.Producer{}

	for _, p := range producers {

		var id int64

		row := ps.Data.DB.QueryRowContext(
			ctx, q,
			p.FirstName,
			StringToNull(p.LastName),
			StringToNull(p.DocumentNumber),
			StringToNull(p.BirthDate),
			StringToNull(p.PhoneNumber),
			StringToNull(p.Address),
			time.Now(),
			time.Now(),
		)

		err := row.Scan(&id)

		if err != nil {
			log.Println(err)
			return pds, err
		}

		p.ID = id

		pds = append(pds, p)
	}

	return pds, nil

}

// GetProducers returns all producers from the database
func (ps *ProducerStorage) GetProducers(ctx context.Context) ([]producer.Producer, error) {

	q := `
	SELECT id, first_name, last_name, document_number, 
		birth_date, phone_number, address, 
		is_active, created_at, updated_at
		FROM producers;
	`

	rows, err := ps.Data.DB.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	pds := []producer.Producer{}

	for rows.Next() {
		p, err := ScanRowProducers(rows)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		pds = append(pds, p)
	}

	return pds, nil
}

// GetProducerByID returns a producer from the database by id
func (ps *ProducerStorage) GetProducerByID(ctx context.Context, id string) (producer.Producer, error) {

	q := `
	SELECT id, first_name, last_name, document_number,
		birth_date, phone_number, address,
		is_active, created_at, updated_at
		FROM producers
		WHERE id = $1;
	`

	row := ps.Data.DB.QueryRowContext(ctx, q, id)

	p, err := ScanRowProducers(row)

	if err != nil {
		log.Println(err)
		return p, err
	}

	return p, nil
}

// UpdateProducer updates a producer in the database
func (ps *ProducerStorage) UpdateProducer(ctx context.Context, id string, p producer.Producer) (producer.Producer, error) {

	q := `
	UPDATE producers
		SET first_name = $1, last_name = $2, document_number = $3, 
		birth_date = $4, phone_number = $5, address = $6, updated_at = $7
		WHERE id = $8
		RETURNING id, first_name, last_name, document_number, birth_date, phone_number, address, is_active, created_at, updated_at;
	`

	row := ps.Data.DB.QueryRowContext(
		ctx, q,
		p.FirstName,
		StringToNull(p.LastName),
		StringToNull(p.DocumentNumber),
		StringToNull(p.BirthDate),
		StringToNull(p.PhoneNumber),
		StringToNull(p.Address),
		time.Now(),
		id,
	)

	producer, err := ScanRowProducers(row)

	if err != nil {
		log.Println(err)
		return producer, err
	}

	return producer, nil
}

// PartialUpdateProducer updates a producer in the database
func (ps *ProducerStorage) PartialUpdateProducer(ctx context.Context, id string, p producer.Producer) (producer.Producer, error) {

	q := `
	UPDATE producers
		SET
			first_name = CASE WHEN $1 = '' THEN first_name ELSE $1 END,
			last_name = CASE WHEN $2 = '' THEN last_name ELSE $2 END,
			document_number = CASE WHEN $3 = '' THEN document_number ELSE $3 END,
			birth_date = CASE WHEN $4 = '' THEN birth_date ELSE $4 END,
			phone_number = CASE WHEN $5 = '' THEN phone_number ELSE $5 END,
			address = CASE WHEN $6 = '' THEN address ELSE $6 END,
			is_active = 
				CASE
					WHEN $7 = TRUE AND is_active = FALSE THEN TRUE
					WHEN $7 = FALSE AND is_active = TRUE THEN FALSE
					WHEN $7 = NULL THEN is_active
					ELSE is_active
				END,
			updated_at = $8
		WHERE id = $9
		RETURNING id, first_name, last_name, document_number, 
			birth_date, phone_number, address, is_active, 
			created_at, updated_at;
	`

	row := ps.Data.DB.QueryRowContext(
		ctx, q,
		p.FirstName,
		p.LastName,
		p.DocumentNumber,
		p.BirthDate,
		p.PhoneNumber,
		p.Address,
		p.IsActive,
		time.Now(),
		id,
	)

	producer, err := ScanRowProducers(row)

	if err != nil {
		log.Println(err)
		return producer, err
	}

	return producer, nil
}

// DeleteProducer deletes a producer from the database
func (ps *ProducerStorage) DeleteProducer(ctx context.Context, id string) (producer.Producer, error) {

	q := `
	DELETE FROM producers
		WHERE id = $1
		RETURNING id, first_name, last_name, document_number, birth_date, phone_number, address, is_active, created_at, updated_at;
	`

	row := ps.Data.DB.QueryRowContext(
		ctx, q,
		id,
	)

	producer, err := ScanRowProducers(row)

	if err != nil {
		log.Println(err)
		return producer, err
	}

	return producer, nil
}
