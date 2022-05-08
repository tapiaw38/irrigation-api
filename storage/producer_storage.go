package storage

import (
	"context"
	"log"

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

		var id uint

		row := ps.Data.DB.QueryRowContext(ctx, q, p.FirstName, p.LastName, p.DocumentNumber, p.BirthDate, p.PhoneNumber, p.Address, p.CreatedAt, p.UpdatedAt)

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
	SELECT id, first_name, last_name, document_number, birth_date, phone_number, address, created_at, updated_at
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
