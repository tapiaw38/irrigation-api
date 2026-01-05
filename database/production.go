package database

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
	"github.com/tapiaw38/irrigation-api/utils"
)

// CreateProductions creates a new production in the database
func (pd *PostgresRepository) CreateProductions(ctx context.Context, productions []models.Production) ([]models.Production, error) {

	q := `
	INSERT INTO productions (
		producer, lote_number, entry, name, production_type, area,
		cultivated_area, area_coordinates, cultivated_area_coordinates,
		picture, cadastral_registration, district, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id;
	`

	pdr := []models.Production{}

	for _, p := range productions {

		var id int64

		area, err := utils.CalculatePolygonArea(p.AreaCoordinates)
		if err != nil {
			log.Println(err)
			return pdr, err
		}
		p.Area = float64(int(area*100)) / 100

		cultivatedArea, err := utils.CalculatePolygonArea(p.CultivatedAreaCoordinates)
		if err != nil {
			log.Println(err)
			return pdr, err
		}
		p.CultivatedArea = float64(int(cultivatedArea*100)) / 100

		row := pd.db.QueryRowContext(
			ctx, q,
			p.Producer,
			StringToNull(p.LoteNumber),
			StringToNull(p.Entry),
			p.Name,
			p.ProductionType,
			FloatToNull(p.Area),
			FloatToNull(p.CultivatedArea),
			StringToNull(p.AreaCoordinates),
			StringToNull(p.CultivatedAreaCoordinates),
			StringToNull(p.Picture),
			StringToNull(p.CadastralRegistration),
			StringToNull(p.District),
			time.Now(),
			time.Now(),
		)

		err = row.Scan(&id)

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
func (pd *PostgresRepository) GetProductions(ctx context.Context) ([]models.ProductionResponse, error) {

	q := `
	SELECT productions.id, producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, productions.name,
		productions.production_type, productions.area, productions.cultivated_area,
		productions.area_coordinates, productions.cultivated_area_coordinates,
		productions.picture, productions.cadastral_registration, productions.district,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id;
	`

	rows, err := pd.db.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	productions := []models.ProductionResponse{}

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
func (pd *PostgresRepository) GetProductionsByID(ctx context.Context, id string) (models.ProductionResponse, error) {

	q := `
	SELECT productions.id, producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, productions.name,
		productions.production_type, productions.area, productions.cultivated_area,
		productions.area_coordinates, productions.cultivated_area_coordinates,
		productions.picture, productions.cadastral_registration, productions.district,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		WHERE productions.id = $1;
	`

	row := pd.db.QueryRowContext(ctx, q, id)

	pds, err := ScanRowProductionResponse(row)

	if err != nil {
		log.Println(err)
		return pds, err
	}

	return pds, nil
}

// UpdateProduction updates a production in the database
func (pd *PostgresRepository) UpdateProduction(ctx context.Context, id string, p models.Production) (models.ProductionResponse, error) {

	area, err := utils.CalculatePolygonArea(p.AreaCoordinates)
	if err != nil {
		log.Println(err)
		return models.ProductionResponse{}, err
	}
	p.Area = float64(int(area*100)) / 100

	cultivatedArea, err := utils.CalculatePolygonArea(p.CultivatedAreaCoordinates)
	if err != nil {
		log.Println(err)
		return models.ProductionResponse{}, err
	}
	p.CultivatedArea = float64(int(cultivatedArea*100)) / 100

	q := `
	WITH updated AS (
		UPDATE productions
		SET producer = $1, lote_number = $2, entry = $3,
			name = $4, production_type = $5, area = $6,
			cultivated_area = $7, area_coordinates = $8,
			cultivated_area_coordinates = $9, picture = $10,
			cadastral_registration = $11, district = $12, updated_at = $13
		WHERE id = $14
		RETURNING id, producer, lote_number, entry, name,
			production_type, area, cultivated_area, area_coordinates,
			cultivated_area_coordinates, picture, cadastral_registration,
			district, created_at, updated_at
	)
	SELECT updated.id, producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		updated.lote_number, updated.entry, updated.name,
		updated.production_type, updated.area, updated.cultivated_area,
		updated.area_coordinates, updated.cultivated_area_coordinates,
		updated.picture, updated.cadastral_registration, updated.district,
		updated.created_at, updated.updated_at
	FROM updated
	LEFT JOIN producers ON updated.producer = producers.id
`

	row := pd.db.QueryRowContext(
		ctx, q,
		p.Producer,
		p.LoteNumber,
		p.Entry,
		p.Name,
		p.ProductionType,
		p.Area,
		p.CultivatedArea,
		p.AreaCoordinates,
		p.CultivatedAreaCoordinates,
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
func (pd *PostgresRepository) PartialUpdateProduction(ctx context.Context, id string, p models.Production) (models.ProductionResponse, error) {

	area := p.Area
	if p.AreaCoordinates != "" {
		calculatedArea, err := utils.CalculatePolygonArea(p.AreaCoordinates)
		if err != nil {
			log.Println(err)
			return models.ProductionResponse{}, err
		}
		area = float64(int(calculatedArea*100)) / 100
	}

	cultivatedArea := p.CultivatedArea
	if p.CultivatedAreaCoordinates != "" {
		calculatedCultivatedArea, err := utils.CalculatePolygonArea(p.CultivatedAreaCoordinates)
		if err != nil {
			log.Println(err)
			return models.ProductionResponse{}, err
		}
		cultivatedArea = float64(int(calculatedCultivatedArea*100)) / 100
	}

	q := `
	WITH updated AS (
		UPDATE productions
		SET
			producer = CASE WHEN $1 = 0 THEN producer ELSE $1 END,
			lote_number = CASE WHEN $2 = '' THEN lote_number ELSE $2 END,
			entry = CASE WHEN $3 = '' THEN entry ELSE $3 END,
			name = CASE WHEN $4 = '' THEN name ELSE $4 END,
			production_type = CASE WHEN $5 = '' THEN production_type ELSE $5 END,
			area = CASE WHEN $6::numeric = 0.0 THEN area ELSE $6::numeric END,
			cultivated_area = CASE WHEN $7::numeric = 0.0 THEN cultivated_area ELSE $7::numeric END,
			area_coordinates = CASE WHEN $8::text = '' THEN area_coordinates ELSE $8::jsonb END,
			cultivated_area_coordinates = CASE WHEN $9::text = '' THEN cultivated_area_coordinates ELSE $9::jsonb END,
			picture = CASE WHEN $10 = '' THEN picture ELSE $10 END,
			cadastral_registration = CASE WHEN $11 = '' THEN cadastral_registration ELSE $11 END,
			district = CASE WHEN $12 = '' THEN district ELSE $12 END,
			updated_at = $13
		WHERE id = $14
		RETURNING id, producer, lote_number, entry, name,
			production_type, area, cultivated_area, area_coordinates,
			cultivated_area_coordinates, picture, cadastral_registration,
			district, created_at, updated_at
	)
	SELECT updated.id, producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		updated.lote_number, updated.entry, updated.name,
		updated.production_type, updated.area, updated.cultivated_area,
		updated.area_coordinates, updated.cultivated_area_coordinates,
		updated.picture, updated.cadastral_registration, updated.district,
		updated.created_at, updated.updated_at
	FROM updated
	LEFT JOIN producers ON updated.producer = producers.id
`

	row := pd.db.QueryRowContext(
		ctx, q,
		p.Producer,
		p.LoteNumber,
		p.Entry,
		p.Name,
		p.ProductionType,
		area,
		cultivatedArea,
		p.AreaCoordinates,
		p.CultivatedAreaCoordinates,
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
func (pd *PostgresRepository) DeleteProduction(ctx context.Context, id string) (models.Production, error) {

	q := `
	DELETE FROM productions
		WHERE id = $1
		RETURNING id, producer, lote_number, entry, name,
			production_type, area, cultivated_area, area_coordinates,
			cultivated_area_coordinates, picture, cadastral_registration,
			district, created_at, updated_at;
	`

	row := pd.db.QueryRowContext(ctx, q, id)

	pds, err := ScanRowProduction(row)

	if err != nil {
		log.Println(err)
		return pds, err
	}

	return pds, nil
}
