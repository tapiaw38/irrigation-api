package storage

import (
	"context"
	"time"

	"github.com/tapiaw38/irrigation-api/models/user"
)

type UserStorage struct {
	Data *Data
}

// CheckUser checks if a user and email exists in the database
func (ur *UserStorage) CheckUser(email string) (user.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	q := `
	SELECT id, first_name, last_name, username, email, password, picture, is_active, created_at, updated_at
		FROM users
		WHERE email = $1;
	`

	rows := ur.Data.DB.QueryRowContext(ctx, q, email)

	var user user.User

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Picture,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, false
	}

	return user, true

}

// CreateUser inserts a new user into the database
func (ur *UserStorage) CreateUser(ctx context.Context, u *user.User) (user.User, error) {

	var user user.User

	q := `
    INSERT INTO users (first_name, last_name, username, email, picture, password, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, first_name, last_name, username, email, picture, is_active, created_at, updated_at;
    `

	if err := u.HashPassword(); err != nil {
		return user, err
	}

	row := ur.Data.DB.QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.PasswordHash, u.IsActive, time.Now(), time.Now(),
	)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Picture,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}

// DeleteUser deletes a user from the database
func (ur *UserStorage) DeleteUser(ctx context.Context, id string) error {

	q := `
	UPDATE users
		SET is_active = false
		WHERE id = $1;
	`

	rows, err := ur.Data.DB.PrepareContext(ctx, q)

	if err != nil {
		return err
	}

	defer rows.Close()

	_, err = rows.ExecContext(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

// Get all users from database
func (ur *UserStorage) GetUsers(ctx context.Context) ([]user.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, password, is_active, created_at, updated_at
		FROM users;
	`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user.User{}

	for rows.Next() {
		var u user.User

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Email,
			&u.Picture,
			&u.PasswordHash,
			&u.IsActive,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}

// Get user by id from database
func (ur *UserStorage) GetUserById(ctx context.Context, id string) (user.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, is_active, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	row := ur.Data.DB.QueryRowContext(
		ctx, q, id,
	)

	u := user.User{}

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Picture,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil

}

// Get user by username from database
func (ur *UserStorage) GetUserByUsername(ctx context.Context, username string) (user.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, is_active, created_at, updated_at
		FROM users
		WHERE username = $1;
	`

	row := ur.Data.DB.QueryRowContext(
		ctx, q, username,
	)

	u := user.User{}

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Picture,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil

}

// UpdateUser updates a user in the database
func (ur *UserStorage) UpdateUser(ctx context.Context, id string, u user.User) (user.User, error) {

	var user user.User

	q := `
	UPDATE users
		SET first_name = $1, last_name = $2, username = $3, email = $4, picture = $5, is_active = $6, updated_at = $7
		WHERE id = $8
		RETURNING id, first_name, last_name, username, email, picture, is_active, created_at, updated_at;
	`

	rows := ur.Data.DB.QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.IsActive, time.Now(), id,
	)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Picture,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}
