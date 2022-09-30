package database

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

// CreateTurn creates a new Turn.
func (ts *PostgresRepository) CreateTurn(ctx context.Context, turn models.Turn) (models.Turn, error) {
	q := `
	INSERT INTO turns (start_date, turn_hours, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, start_date, turn_hours, end_date, created_at, updated_at;
	`
	row := ts.db.QueryRowContext(
		ctx, q,
		turn.StartDate,
		turn.TurnHours,
		StringToNull(turn.EndDate),
		time.Now(),
		time.Now(),
	)
	t, err := ScanRowTurn(row)

	if err != nil {
		log.Println(err)
		return turn, err
	}

	return t, nil
}

// GetTurns returns all the Turns.
func (ts *PostgresRepository) GetTurns(ctx context.Context) ([]models.TurnResponse, error) {
	q := `
	SELECT id, start_date, turn_hours, end_date, created_at, updated_at
		FROM turns
		ORDER BY start_date DESC;
	`
	rows, err := ts.db.QueryContext(ctx, q)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	turns := []models.TurnResponse{}

	for rows.Next() {
		trs, err := ScanRowTurnResponse(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		q := `
		SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
			producers.document_number, producers.birth_date, producers.phone_number,
			producers.address,
			productions.lote_number, productions.entry, 
			productions.name, productions.production_type, productions.area, 
			productions.cultivated_area, productions.latitude, productions.longitude, 
			productions.picture, productions.cadastral_registration, productions.district,
			intakes.id, intakes.intake_number,
			intakes_productions.watering_order,
			productions.created_at, productions.updated_at
			FROM productions
			LEFT JOIN producers ON productions.producer = producers.id
			LEFT JOIN intakes_productions
            ON intakes_productions.production_id=productions.id
			LEFT JOIN intakes ON intakes.id = intakes_productions.intake_id
			LEFT JOIN turns_productions
			ON turns_productions.production_id = productions.id
			WHERE turns_productions.turn_id = $1
			ORDER BY intakes_productions.intake_id ASC, intakes_productions.watering_order ASC;
		`
		rows, err := ts.db.QueryContext(
			ctx, q,
			trs.ID,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		for rows.Next() {
			pds, err := ScanRowProductionTurnResponse(rows)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			pds.WateringHour = 2 * pds.Area
		}

		turns = append(turns, trs)
	}

	return turns, nil
}

// UpdateTurn updates a Turn.
func (ts *PostgresRepository) UpdateTurn(ctx context.Context, id string, turn models.Turn) (models.TurnResponse, error) {
	q := `
	UPDATE turns
		SET start_date = $1, updated_at = $2
		WHERE id = $3
		RETURNING id, start_date, turn_hours, end_date, created_at, updated_at;
	`

	row := ts.db.QueryRowContext(
		ctx, q,
		turn.StartDate,
		time.Now(),
		id,
	)

	t, err := ScanRowTurnResponse(row)

	if err != nil {
		log.Println(err)
		return t, err
	}

	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes.id, intakes.intake_number,
		intakes_productions.watering_order,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		LEFT JOIN intakes ON intakes.id = intakes_productions.intake_id
		LEFT JOIN turns_productions
		ON turns_productions.production_id = productions.id
		WHERE turns_productions.turn_id = $1
		ORDER BY intakes_productions.intake_id ASC, intakes_productions.watering_order ASC;
		`
	rows, err := ts.db.QueryContext(
		ctx, q,
		t.ID,
	)

	if err != nil {
		log.Println(err)
		return t, err
	}

	for rows.Next() {
		pds, err := ScanRowProductionTurnResponse(rows)
		if err != nil {
			log.Println(err)
			return t, err
		}
		pds.WateringHour = 2 * pds.Area
		t.Productions = append(t.Productions, pds)
	}

	return t, nil
}

// DeleteTurn deletes a Turn.
func (ts *PostgresRepository) DeleteTurn(ctx context.Context, id string) (models.Turn, error) {
	q := `
	DELETE FROM turns
		WHERE id = $1
		RETURNING id, start_date, turn_hours, end_date, created_at, updated_at;
	`

	row := ts.db.QueryRowContext(
		ctx, q,
		id,
	)

	t, err := ScanRowTurn(row)

	if err != nil {
		log.Println(err)
		return t, err
	}

	return t, nil
}

// GetTurnByID returns a Turn by ID.
func (ts *PostgresRepository) GetTurnByID(ctx context.Context, id string) (models.TurnResponse, error) {
	q := `
	SELECT id, start_date, turn_hours, end_date, created_at, updated_at
		FROM turns
		WHERE id = $1;
	`

	row := ts.db.QueryRowContext(ctx, q, id)

	turn, err := ScanRowTurnResponse(row)

	if err != nil {
		log.Println(err)
		return turn, err
	}

	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes.id, intakes.intake_number,
		intakes_productions.watering_order,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		LEFT JOIN intakes ON intakes.id = intakes_productions.intake_id
		LEFT JOIN turns_productions
		ON turns_productions.production_id = productions.id
		WHERE turns_productions.turn_id = $1
		ORDER BY intakes_productions.intake_id ASC, intakes_productions.watering_order ASC;
		`

	rows, err := ts.db.QueryContext(
		ctx, q,
		id,
	)

	if err != nil {
		log.Println(err)
		return turn, err
	}

	for rows.Next() {
		pds, err := ScanRowProductionTurnResponse(rows)
		if err != nil {
			log.Println(err)
			return turn, err
		}
		pds.WateringHour = 2 * pds.Area
		turn.Productions = append(turn.Productions, pds)
	}

	return turn, nil
}

// CreateTurnProduction creates a TurnProduction.
func (ts *PostgresRepository) CreateTurnProduction(ctx context.Context, turnID string, turnProduction models.TurnProduction) (models.TurnResponse, error) {
	var tps models.TurnResponse
	q := `
	INSERT INTO turns_productions (turn_id, production_id)
		VALUES ($1, $2)
		RETURNING turn_id, production_id;
	`
	row := ts.db.QueryRowContext(
		ctx, q,
		turnID,
		turnProduction.ProductionID,
	)
	tp, err := ScanRowTurnProduction(row)
	if err != nil {
		log.Println(err)
		return tps, err
	}
	q = `
	SELECT id, start_date, turn_hours, end_date, created_at, updated_at
		FROM turns
		WHERE id = $1;
	`
	row = ts.db.QueryRowContext(ctx, q, tp.TurnID)
	tps, err = ScanRowTurnResponse(row)
	if err != nil {
		log.Println(err)
		return tps, err
	}
	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes.id, intakes.intake_number,
		intakes_productions.watering_order,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		LEFT JOIN intakes ON intakes.id = intakes_productions.intake_id
		LEFT JOIN turns_productions
		ON turns_productions.production_id = productions.id
		WHERE turns_productions.turn_id = $1
		ORDER BY intakes_productions.intake_id ASC, intakes_productions.watering_order ASC;
	`
	rows, err := ts.db.QueryContext(
		ctx, q,
		tps.ID,
	)
	if err != nil {
		log.Println(err)
		return tps, err
	}
	for rows.Next() {
		pds, err := ScanRowProductionTurnResponse(rows)
		if err != nil {
			log.Println(err)
			return tps, err
		}
		pds.WateringHour = 2 * pds.Area
		tps.Productions = append(tps.Productions, pds)
	}
	return tps, nil
}

// DeleteTurnProduction deletes a TurnProduction.
func (ts *PostgresRepository) DeleteTurnProduction(ctx context.Context, turnID string, turnProduction models.TurnProduction) (models.TurnResponse, error) {
	var tps models.TurnResponse
	q := `
	DELETE FROM turns_productions
		WHERE turn_id = $1 AND production_id = $2
		RETURNING turn_id, production_id;
	`
	row := ts.db.QueryRowContext(
		ctx, q,
		turnID,
		turnProduction.ProductionID,
	)
	tp, err := ScanRowTurnProduction(row)
	if err != nil {
		log.Println(err)
		return tps, err
	}
	q = `
	SELECT id, start_date, turn_hours, end_date, created_at, updated_at
		FROM turns
		WHERE id = $1;
	`
	row = ts.db.QueryRowContext(ctx, q, tp.TurnID)
	tps, err = ScanRowTurnResponse(row)
	if err != nil {
		log.Println(err)
		return tps, err
	}
	q = `
	SELECT productions.id,  producers.id, producers.first_name, producers.last_name,
		producers.document_number, producers.birth_date, producers.phone_number,
		producers.address,
		productions.lote_number, productions.entry, 
		productions.name, productions.production_type, productions.area, 
		productions.cultivated_area, productions.latitude, productions.longitude, 
		productions.picture, productions.cadastral_registration, productions.district,
		intakes.id, intakes.intake_number,
		intakes_productions.watering_order,
		productions.created_at, productions.updated_at
		FROM productions
		LEFT JOIN producers ON productions.producer = producers.id
		LEFT JOIN intakes_productions
		ON intakes_productions.production_id=productions.id
		LEFT JOIN intakes ON intakes.id = intakes_productions.intake_id
		LEFT JOIN turns_productions
		ON turns_productions.production_id = productions.id
		WHERE turns_productions.turn_id = $1
		ORDER BY intakes_productions.intake_id ASC, intakes_productions.watering_order ASC;
	`
	rows, err := ts.db.QueryContext(
		ctx, q,
		tps.ID,
	)
	if err != nil {
		log.Println(err)
		return tps, err
	}
	for rows.Next() {
		pds, err := ScanRowProductionTurnResponse(rows)
		if err != nil {
			log.Println(err)
			return tps, err
		}
		pds.WateringHour = 2 * pds.Area
		tps.Productions = append(tps.Productions, pds)
	}
	return tps, nil
}
