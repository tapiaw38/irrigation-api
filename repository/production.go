package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Repository handle the CRUD operations with Productions.

func CreateProductions(ctx context.Context, productions []models.Production) ([]models.Production, error) {
	return implementation.CreateProductions(ctx, productions)
}

func GetProductions(ctx context.Context) ([]models.ProductionResponse, error) {
	return implementation.GetProductions(ctx)
}

func GetProductionsByID(ctx context.Context, id string) (models.ProductionResponse, error) {
	return implementation.GetProductionsByID(ctx, id)
}

func UpdateProduction(ctx context.Context, id string, production models.Production) (models.ProductionResponse, error) {
	return implementation.UpdateProduction(ctx, id, production)
}

func PartialUpdateProduction(ctx context.Context, id string, production models.Production) (models.ProductionResponse, error) {
	return implementation.PartialUpdateProduction(ctx, id, production)
}

func DeleteProduction(ctx context.Context, id string) (models.Production, error) {
	return implementation.DeleteProduction(ctx, id)
}
