package intake

import "context"

// Storage handle the CRUD operations with Intakes.
type Storage interface {
	CreateIntakes(ctx context.Context, intakes []Intake) ([]Intake, error)
	GetIntakes(ctx context.Context) ([]IntakeResponse, error)
	GetIntakeByID(ctx context.Context, id string) (IntakeResponse, error)
	UpdateIntake(ctx context.Context, id string, intake Intake) (IntakeResponse, error)
	DeleteIntake(ctx context.Context, id string) (Intake, error)
	CreateIntakeProduction(ctx context.Context, intakeID string, intakeProduction IntakeProduction) (IntakeResponse, error)
	UpdateIntakeProduction(ctx context.Context, intakeID string, intakeProduction IntakeProduction) (IntakeResponse, error)
	DeleteIntakeProduction(ctx context.Context, intakeID string, intakeProduction IntakeProduction) (IntakeResponse, error)
}
