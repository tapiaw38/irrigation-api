package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Configuration operations

func GetConfiguration(ctx context.Context) (models.Configuration, error) {
	return implementation.GetConfiguration(ctx)
}

func UpdateConfiguration(ctx context.Context, config *models.Configuration) (models.Configuration, error) {
	return implementation.UpdateConfiguration(ctx, config)
}
