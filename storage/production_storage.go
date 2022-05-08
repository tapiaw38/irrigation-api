package storage

import (
	"context"
	"log"

	"github.com/tapiaw38/irrigation-api/models/production"
)

type ProductionStorage struct {
	Data *Data
}

// CreateProductions creates a new production in the database
func (pd *ProductionStorage) CreateProductions(ctx context.Context, productions []production.Production) ([]production.Production, error) {

	q := `
	INSERT INTO productions (producer, lote_number, entry, name, production_type, latitude, longitude, picture, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;
	`

	pdr := []production.Production{}

	for _, p := range productions {

		var id uint

		row := pd.Data.DB.QueryRowContext(ctx, q, p.Producer, p.LoteNumber, p.Entry, p.Name, p.ProductionType, p.Latitude, p.Longitude, p.Picture, p.CreatedAt, p.UpdatedAt)

		err := row.Scan(&id)

		if err != nil {
			log.Println(err)
			return pdr, err
		}

		p.ID = id

		pdr = append(pdr, p)
	}

	return pdr, nil
}

func (pd *ProductionStorage) GetProductions(ctx context.Context) ([]production.Production, error) {

	q := `
	SELECT id, producer, lote_number, entry, name, production_type, latitude, longitude, picture, created_at, updated_at
		FROM productions;
	`

	rows, err := pd.Data.DB.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	pdrs := []production.Production{}

	for rows.Next() {
		pdr, err := ScanRowProduction(rows)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		pdrs = append(pdrs, pdr)
	}

	return pdrs, nil
}
