package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// repository handle the CRUD operations.
type Repository interface {
	// user
	CheckUser(email string) (models.User, bool)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	UpdateUser(ctx context.Context, id string, user models.User) (models.User, error)
	PartialUpdateUser(ctx context.Context, id string, user models.User) (models.User, error)
	// producer
	CreateProducers(ctx context.Context, producers []models.Producer) ([]models.Producer, error)
	GetProducers(ctx context.Context) ([]models.Producer, error)
	GetProducerByID(ctx context.Context, id string) (models.Producer, error)
	UpdateProducer(ctx context.Context, id string, producer models.Producer) (models.Producer, error)
	PartialUpdateProducer(ctx context.Context, id string, producer models.Producer) (models.Producer, error)
	DeleteProducer(ctx context.Context, id string) (models.Producer, error)
	// production
	CreateProductions(ctx context.Context, productions []models.Production) ([]models.Production, error)
	GetProductions(ctx context.Context) ([]models.ProductionResponse, error)
	GetProductionsByID(ctx context.Context, id string) (models.ProductionResponse, error)
	UpdateProduction(ctx context.Context, id string, production models.Production) (models.ProductionResponse, error)
	PartialUpdateProduction(ctx context.Context, id string, production models.Production) (models.ProductionResponse, error)
	DeleteProduction(ctx context.Context, id string) (models.Production, error)
	// sections
	CreateSections(ctx context.Context, sections []models.Section) ([]models.Section, error)
	GetSections(ctx context.Context) ([]models.Section, error)
	GetSectionByID(ctx context.Context, id string) (models.Section, error)
	UpdateSection(ctx context.Context, id string, section models.Section) (models.Section, error)
	DeleteSection(ctx context.Context, id string) (models.Section, error)
	// intakes
	CreateIntakes(ctx context.Context, intakes []models.Intake) ([]models.Intake, error)
	GetIntakes(ctx context.Context) ([]models.IntakeResponse, error)
	GetIntakeByID(ctx context.Context, id string) (models.IntakeResponse, error)
	UpdateIntake(ctx context.Context, id string, intake models.Intake) (models.IntakeResponse, error)
	DeleteIntake(ctx context.Context, id string) (models.Intake, error)
	CreateIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error)
	UpdateIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error)
	DeleteIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error)
	// turns
	CreateTurn(ctx context.Context, turn models.Turn) (models.Turn, error)
	GetTurns(ctx context.Context) ([]models.TurnResponse, error)
	GetTurnByID(ctx context.Context, id string) (models.TurnResponse, error)
	UpdateTurn(ctx context.Context, id string, turn models.Turn) (models.TurnResponse, error)
	DeleteTurn(ctx context.Context, id string) (models.Turn, error)
	CreateTurnProduction(ctx context.Context, turnID string, turnProduction models.TurnProduction) (models.TurnResponse, error)
	DeleteTurnProduction(ctx context.Context, turnID string, turnProduction models.TurnProduction) (models.TurnResponse, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}
