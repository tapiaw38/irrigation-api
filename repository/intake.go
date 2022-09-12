package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Storage handle the CRUD operations with Intakes.

func CreateIntakes(ctx context.Context, intakes []models.Intake) ([]models.Intake, error) {
	return implementation.CreateIntakes(ctx, intakes)
}

func GetIntakes(ctx context.Context) ([]models.IntakeResponse, error) {
	return implementation.GetIntakes(ctx)
}

func GetIntakeByID(ctx context.Context, id string) (models.IntakeResponse, error) {
	return implementation.GetIntakeByID(ctx, id)
}

func UpdateIntake(ctx context.Context, id string, intake models.Intake) (models.IntakeResponse, error) {
	return implementation.UpdateIntake(ctx, id, intake)
}

func DeleteIntake(ctx context.Context, id string) (models.Intake, error) {
	return implementation.DeleteIntake(ctx, id)
}

func CreateIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error) {
	return implementation.CreateIntakeProduction(ctx, intakeID, intakeProduction)
}

func UpdateIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error) {
	return implementation.UpdateIntakeProduction(ctx, intakeID, intakeProduction)
}
func DeleteIntakeProduction(ctx context.Context, intakeID string, intakeProduction models.IntakeProduction) (models.IntakeResponse, error) {
	return implementation.DeleteIntakeProduction(ctx, intakeID, intakeProduction)
}
