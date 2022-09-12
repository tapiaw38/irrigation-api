package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Storage handle the CRUD operations with Sections.

func CreateSections(ctx context.Context, sections []models.Section) ([]models.Section, error) {
	return implementation.CreateSections(ctx, sections)
}

func GetSections(ctx context.Context) ([]models.Section, error) {
	return implementation.GetSections(ctx)
}

func GetSectionByID(ctx context.Context, id string) (models.Section, error) {
	return implementation.GetSectionByID(ctx, id)
}

func UpdateSection(ctx context.Context, id string, section models.Section) (models.Section, error) {
	return implementation.UpdateSection(ctx, id, section)
}

func DeleteSection(ctx context.Context, id string) (models.Section, error) {
	return implementation.DeleteSection(ctx, id)
}
