package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Storage handle the CRUD operations with Turns.

func CreateTurn(ctx context.Context, turn models.Turn) (models.Turn, error) {
	return implementation.CreateTurn(ctx, turn)
}

func GetTurns(ctx context.Context) ([]models.TurnResponse, error) {
	return implementation.GetTurns(ctx)
}

func GetTurnByID(ctx context.Context, id string) (models.TurnResponse, error) {
	return implementation.GetTurnByID(ctx, id)
}

func UpdateTurn(ctx context.Context, id string, turn models.Turn) (models.TurnResponse, error) {
	return implementation.UpdateTurn(ctx, id, turn)
}

func DeleteTurn(ctx context.Context, id string) (models.Turn, error) {
	return implementation.DeleteTurn(ctx, id)
}
func CreateTurnProduction(ctx context.Context, turnID string, turnProduction models.TurnProduction) (models.TurnResponse, error) {
	return implementation.CreateTurnProduction(ctx, turnID, turnProduction)
}

func DeleteTurnProduction(ctx context.Context, turnID string, turnProduction models.TurnProduction) (models.TurnResponse, error) {
	return implementation.DeleteTurnProduction(ctx, turnID, turnProduction)
}
