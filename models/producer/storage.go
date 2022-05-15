package producer

import "context"

// Storage handle the CRUD operations with Producers.
type Storage interface {
	CreateProducers(ctx context.Context, producers []Producer) ([]Producer, error)
	GetProducers(ctx context.Context) ([]Producer, error)
	GetProducerByID(ctx context.Context, id string) (Producer, error)
	UpdateProducer(ctx context.Context, id string, producer Producer) (Producer, error)
	PartialUpdateProducer(ctx context.Context, id string, producer Producer) (Producer, error)
}
