package production

import "context"

// Storage handle the CRUD operations with Productions.
type Storage interface {
	CreateProductions(ctx context.Context, productions []Production) ([]Production, error)
	GetProductions(ctx context.Context) ([]Production, error)
}
