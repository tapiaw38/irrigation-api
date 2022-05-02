package producer

import "context"

// Storage handle the CRUD operations with Producers.
type Storage interface {
	CreateProducers(ctx context.Context, producers []Producer) ([]Producer, error)
}
