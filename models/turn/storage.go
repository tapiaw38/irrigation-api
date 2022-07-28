package turn

import "context"

// Storage handle the CRUD operations with Turns.
type Storage interface {
	CreateTurn(ctx context.Context, turn Turn) (Turn, error)
	GetTurns(ctx context.Context) ([]TurnResponse, error)
	GetTurnByID(ctx context.Context, id string) (TurnResponse, error)
	// UpdateTurn(ctx context.Context, id string, turn Turn) (Turn, error)
	// DeleteTurn(ctx context.Context, id string) (Turn, error)
}
