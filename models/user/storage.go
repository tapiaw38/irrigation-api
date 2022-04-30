package user

import "context"

// Storage handle the CRUD operations with Users.
type Storage interface {
	CheckUser(email string) (User, bool)
	CreateUser(ctx context.Context, user *User) (User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context) ([]User, error)
	GetUserById(ctx context.Context, id string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateUser(ctx context.Context, id string, user User) (User, error)
	PartialUpdateUser(ctx context.Context, id string, user User) (User, error)
}
