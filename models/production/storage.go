package production

import "context"

// Storage handle the CRUD operations with Productions.
type Storage interface {
	CreateProductions(ctx context.Context, productions []Production) ([]Production, error)
	GetProductions(ctx context.Context) ([]ProductionResponse, error)
	GetProductionsByID(ctx context.Context, id string) (ProductionResponse, error)
	UpdateProduction(ctx context.Context, id string, production Production) (ProductionResponse, error)
	DeleteProduction(ctx context.Context, id string) (Production, error)
}
