package storage

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models/production"
)

type ProductionStorage struct {
	Data *Data
}

// CreateProductions creates a new production in the database
func (pd *ProductionStorage) CreateProductions(ctx context.Context, productions []production.Production) ([]production.Production, error) {

	q := `
	INSERT INTO productions (producer, lote_number, entry, name, production_type, area, latitude, longitude, picture, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id;
	`

	pdr := []production.Production{}

	for _, p := range productions {

		var id int64

		row := pd.Data.DB.QueryRowContext(
			ctx, q,
			p.Producer,
			p.LoteNumber,
			p.Entry,
			p.Name,
			p.ProductionType,
			p.Area,
			p.Latitude,
			p.Longitude,
			p.Picture,
			time.Now(),
			time.Now(),
		)

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

// GetProductions returns all productions from the database
func (pd *ProductionStorage) GetProductions(ctx context.Context) ([]production.ProductionResponse, error) {

	q := `
	SELECT productions.id, producers.id, producers.first_name, producers.last_name, 
		producers.document_number, producers.birth_date, producers.phone_number, 
		producers.address,
		productions.lote_number, productions.entry, productions.name, 
		productions.production_type, productions.area, productions.latitude, 
		productions.longitude, productions.picture, productions.created_at, 
		productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id;
	`

	rows, err := pd.Data.DB.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	productions := []production.ProductionResponse{}

	for rows.Next() {
		pds, err := ScanRowProductionResponse(rows)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		productions = append(productions, pds)
	}

	return productions, nil
}
