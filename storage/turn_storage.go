package storage

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models/turn"
)

type TurnStorage struct {
	Data *Data
}

// CreateTurn creates a new Turn.
func (ts *TurnStorage) CreateTurn(ctx context.Context, turn turn.Turn) (turn.Turn, error) {

	q := `
	INSERT INTO turns (start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, start_date, end_date, created_at, updated_at;
	`

	row := ts.Data.DB.QueryRowContext(
		ctx, q,
		turn.StartDate,
		turn.EndDate,
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
