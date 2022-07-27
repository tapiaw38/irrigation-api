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
	INSERT INTO productions (
		producer, lote_number, entry, name, production_type, area, 
		cultivated_area, latitude, longitude, picture, 
		cadastral_registration, district, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id;
	`

	pdr := []production.Production{}

	for _, p := range productions {

		var id int64

		row := pd.Data.DB.QueryRowContext(
			ctx, q,
			p.Producer,
			StringToNull(p.LoteNumber),
			StringToNull(p.Entry),
			p.Name,
			p.ProductionType,
			FloatToNull(p.Area),
			FloatToNull(p.CultivatedArea),
			FloatToNull(p.Latitude),
			FloatToNull(p.Longitude),
			StringToNull(p.Picture),
			StringToNull(p.CadastralRegistration),
			StringToNull(p.District),
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
		productions.production_type, productions.area, productions.cultivated_area, 
		productions.latitude, productions.longitude, productions.picture,
		productions.cadastral_registration, productions.district,
		productions.created_at, productions.updated_at
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

// GetProductionsByID return a production from the database by id
func (pd *ProductionStorage) GetProductionsByID(ctx context.Context, id string) (production.ProductionResponse, error) {

	q := `
	SELECT productions.id, producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, productions.name,
		productions.production_type, productions.area, productions.cultivated_area, 
		productions.latitude, productions.longitude, productions.picture,
		productions.cadastral_registration, productions.district,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		WHERE productions.id = $1;
	`

	row := pd.Data.DB.QueryRowContext(ctx, q, id)

	pds, err := ScanRowProductionResponse(row)

	if err != nil {
		log.Println(err)
		return pds, err
	}

	return pds, nil
}

// UpdateProduction updates a production in the database
func (pd *ProductionStorage) UpdateProduction(ctx context.Context, id string, p production.Production) (production.ProductionResponse, error) {

	q := `
	WITH updated AS (
		UPDATE productions
		SET producer = $1, lote_number = $2, entry = $3, 
			name = $4, production_type = $5, area = $6, 
			cultivated_area = $7, latitude = $8, longitude = $9, 
			picture = $10, cadastral_registration = $11, 
			district = $12, updated_at = $13
		WHERE id = $14
		RETURNING id, producer, lote_number, entry, name, 
			production_type, area, cultivated_area, latitude, longitude, 
			picture, cadastral_registration, district, created_at, updated_at
	)
	SELECT updated.id, producers.id, producers.first_name, producers.last_name, 
		producers.document_number, producers.birth_date, producers.phone_number, 
		producers.address,
		updated.lote_number, updated.entry, updated.name, 
		updated.production_type, updated.area, updated.cultivated_area, 
		updated.latitude, updated.longitude, updated.picture, 
		updated.cadastral_registration, updated.district, 
		updated.created_at, updated.updated_at
	FROM updated
	LEFT JOIN producers ON updated.producer = producers.id
`

	row := pd.Data.DB.QueryRowContext(
		ctx, q,
		p.Producer,
		p.LoteNumber,
		p.Entry,
		p.Name,
		p.ProductionType,
		p.Area,
		p.CultivatedArea,
		p.Latitude,
		p.Longitude,
		p.Picture,
		p.CadastralRegistration,
		p.District,
		time.Now(),
		id,
	)

	pds, err := ScanRowProductionResponse(row)

	if err != nil {
		log.Println(err)
		return pds, err
	}

	return pds, nil
}

// UpdateProduction updates a production in the database
func (pd *ProductionStorage) PartialUpdateProduction(ctx context.Context, id string, p production.Production) (production.ProductionResponse, error) {

	q := `
	WITH updated AS (
		UPDATE productions
		SET 
			producer = CASE WHEN $1 = 0 THEN producer ELSE $1 END, 
			lote_number = CASE WHEN $2 = '' THEN lote_number ELSE $2 END, 
			entry = CASE WHEN $3 = '' THEN entry ELSE $3 END,
			name = CASE WHEN $4 = '' THEN name ELSE $4 END,
			production_type = CASE WHEN $5 = '' THEN production_type ELSE $5 END,
			area = CASE WHEN $6 = 0 THEN area ELSE $6 END,
			cultivated_area = CASE WHEN $7 = 0 THEN cultivated_area ELSE $7 END,
			latitude = CASE WHEN $8 = 0 THEN latitude ELSE $8 END,
			longitude = CASE WHEN $9 = 0 THEN longitude ELSE $9 END,
			picture = CASE WHEN $10 = '' THEN picture ELSE $10 END,
			cadastral_registration = CASE WHEN $11 = '' THEN cadastral_registration ELSE $11 END,
			district = CASE WHEN $12 = '' THEN district ELSE $12 END,
			updated_at = $13
		WHERE id = $14
		RETURNING id, producer, lote_number, entry, name, 
			production_type, area, cultivated_area, latitude, longitude, 
			picture, cadastral_registration, district, created_at, updated_at
	)
	SELECT updated.id, producers.id, producers.first_name, producers.last_name, 
		producers.document_number, producers.birth_date, producers.phone_number, 
		producers.address,
		updated.lote_number, updated.entry, updated.name, 
		updated.production_type, updated.area, updated.cultivated_area, 
		updated.latitude, updated.longitude, updated.picture, 
		updated.cadastral_registration, updated.district, 
		updated.created_at, updated.updated_at
	FROM updated
	LEFT JOIN producers ON updated.producer = producers.id
`

	row := pd.Data.DB.QueryRowContext(
		ctx, q,
		p.Producer,
		p.LoteNumber,
		p.Entry,
		p.Name,
		p.ProductionType,
		p.Area,
		p.CultivatedArea,
		p.Latitude,
		p.Longitude,
		p.Picture,
		p.CadastralRegistration,
		p.District,
		time.Now(),
		id,
	)

	pds, err := ScanRowProductionResponse(row)

	if err != nil {
		log.Println(err)
		return pds, err
	}

	return pds, nil
}

// DeleteProduction deletes a production from the database
func (pd *ProductionStorage) DeleteProduction(ctx context.Context, id string) (production.Production, error) {

	q := `
	DELETE FROM productions
		WHERE id = $1
		RETURNING id, producer, lote_number, entry, name,
			production_type, area, cultivated_area, latitude, longitude, 
			picture, cadastral_registration, district, created_at, updated_at;
	`

	row := pd.Data.DB.QueryRowContext(ctx, q, id)

	pds, err := ScanRowProduction(row)

	if err != nil {
		log.Println(err)
		return pds, err
	}

	return pds, nil
}
