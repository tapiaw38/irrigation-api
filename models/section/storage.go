package section

import "context"

// Storage handle the CRUD operations with Sections.
type Storage interface {
	CreateSections(ctx context.Context, sections []Section) ([]Section, error)
	GetSections(ctx context.Context) ([]Section, error)
	GetSectionByID(ctx context.Context, id string) (Section, error)
	UpdateSection(ctx context.Context, id string, section Section) (Section, error)
	DeleteSection(ctx context.Context, id string) (Section, error)
}
