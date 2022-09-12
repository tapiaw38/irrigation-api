package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Repository handle the CRUD operations with Producers.

func CreateProducers(ctx context.Context, producers []models.Producer) ([]models.Producer, error) {
	return implementation.CreateProducers(ctx, producers)
}

func GetProducers(ctx context.Context) ([]models.Producer, error) {
	return implementation.GetProducers(ctx)
}

func GetProducerByID(ctx context.Context, id string) (models.Producer, error) {
	return implementation.GetProducerByID(ctx, id)
}

func UpdateProducer(ctx context.Context, id string, producer models.Producer) (models.Producer, error) {
	return implementation.UpdateProducer(ctx, id, producer)
}

func PartialUpdateProducer(ctx context.Context, id string, producer models.Producer) (models.Producer, error) {
	return implementation.PartialUpdateProducer(ctx, id, producer)
}

func DeleteProducer(ctx context.Context, id string) (models.Producer, error) {
	return implementation.DeleteProducer(ctx, id)
}
