package repository

import (
	"context"

	"github.com/tapiaw38/irrigation-api/models"
)

// Repository handle the CRUD operations with Users.

func CheckUser(email string) (models.User, bool) {
	return implementation.CheckUser(email)
}

func CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	return implementation.CreateUser(ctx, user)
}

func DeleteUser(ctx context.Context, id string) error {
	return implementation.DeleteUser(ctx, id)
}

func GetUsers(ctx context.Context) ([]models.User, error) {
	return implementation.GetUsers(ctx)
}

func GetUserById(ctx context.Context, id string) (models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	return implementation.GetUserByUsername(ctx, username)
}

func UpdateUser(ctx context.Context, id string, user models.User) (models.User, error) {
	return implementation.UpdateUser(ctx, id, user)
}

func PartialUpdateUser(ctx context.Context, id string, user models.User) (models.User, error) {
	return implementation.PartialUpdateUser(ctx, id, user)
}
