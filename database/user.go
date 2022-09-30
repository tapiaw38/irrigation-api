package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

// CheckUser checks if a user and email exists in the database
func (ur *PostgresRepository) CheckUser(email string) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	q := `
	SELECT id, first_name, last_name, username, email, password, picture, phone_number, address, is_active, is_admin, created_at, updated_at
		FROM users
		WHERE email = $1;
	`
	rows := ur.db.QueryRowContext(ctx, q, email)
	var user models.User
	var picture, phoneNumber, address sql.NullString
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&picture,
		&phoneNumber,
		&address,
		&user.IsActive,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	user.Picture = picture.String
	user.PhoneNumber = phoneNumber.String
	user.Address = address.String
	if err != nil {
		log.Println(err)
		return user, false
	}
	return user, true
}

// CreateUser inserts a new user into the database
func (ur *PostgresRepository) CreateUser(ctx context.Context, u *models.User) (models.User, error) {
	q := `
    INSERT INTO users (first_name, last_name, username, email, picture, phone_number, address, password, is_active, is_admin, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        RETURNING id, first_name, last_name, username, email, picture, phone_number, address, is_active, is_admin, created_at, updated_at;
    `
	if err := u.HashPassword(); err != nil {
		return models.User{}, err
	}
	row := ur.db.QueryRowContext(
		ctx, q,
		u.FirstName,
		StringToNull(u.LastName),
		u.Username,
		u.Email,
		StringToNull(u.Picture),
		StringToNull(u.PhoneNumber),
		StringToNull(u.Address),
		u.PasswordHash,
		u.IsActive,
		u.IsAdmin,
		time.Now(),
		time.Now(),
	)
	users, err := ScanRowUsers(row)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}
	return users, nil
}

// DeleteUser deletes a user from the database
func (ur *PostgresRepository) DeleteUser(ctx context.Context, id string) error {
	q := `
	UPDATE users
		SET is_active = false
		WHERE id = $1;
	`
	rows, err := ur.db.PrepareContext(ctx, q)
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
func (ur *PostgresRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	q := `
	SELECT id, first_name, last_name, username, email, picture, phone_number, address, is_active, is_admin, created_at, updated_at
		FROM users;
	`
	rows, err := ur.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		u, err := ScanRowUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Get user by id from database
func (ur *PostgresRepository) GetUserById(ctx context.Context, id string) (models.User, error) {

	q := `
	SELECT id, first_name, last_name, username, email, picture, phone_number, address, is_active, is_admin, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	row := ur.db.QueryRowContext(
		ctx, q, id,
	)

	u, err := ScanRowUsers(row)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// Get user by username from database
func (ur *PostgresRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	q := `
	SELECT id, first_name, last_name, username, email, picture, phone_number, address, is_active, is_admin, created_at, updated_at
		FROM users
		WHERE username = $1;
	`

	row := ur.db.QueryRowContext(
		ctx, q, username,
	)

	u, err := ScanRowUsers(row)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// UpdateUser updates a user in the database
func (ur *PostgresRepository) UpdateUser(ctx context.Context, id string, u models.User) (models.User, error) {
	q := `
	UPDATE users
		SET 
		first_name = $1, last_name = $2, email = $3, picture = $4, phone_number = $5, address = $6, is_active = $7, is_admin = $8, updated_at = $9
		WHERE id = $10
		RETURNING id, first_name, last_name, username, email, picture, phone_number, address, is_active, is_admin, created_at, updated_at;
	`

	row := ur.db.QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Email,
		StringToNull(u.Picture), StringToNull(u.PhoneNumber), StringToNull(u.Address),
		u.IsActive, u.IsAdmin, time.Now(), id,
	)

	u, err := ScanRowUsers(row)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// PartialUpdateUser updates a user in the database
func (ur *PostgresRepository) PartialUpdateUser(ctx context.Context, id string, u models.User) (models.User, error) {
	q := `
	UPDATE users
		SET 
			first_name = CASE WHEN $1 = '' THEN first_name ELSE $1 END, 
			last_name = CASE WHEN $2 = '' THEN last_name ELSE $2 END, 
			email = CASE WHEN $3 = '' THEN email ELSE $3 END,
			picture = CASE WHEN $4 = '' THEN picture ELSE $4 END,
			phone_number = CASE WHEN $5 = '' THEN phone_number ELSE $5 END,
			address = CASE WHEN $6 = '' THEN address ELSE $6 END,
			is_active = 
				CASE 
					WHEN $7 = TRUE AND is_active = FALSE THEN TRUE 
					WHEN $7 = FALSE AND is_active = TRUE THEN FALSE
					WHEN $7 = NULL THEN is_active
					ELSE is_active
				END,
			is_admin = $8,
			updated_at = $9
		WHERE id = $10
		RETURNING id, first_name, last_name, username, email, picture, phone_number, address, is_active, is_admin, created_at, updated_at;
	`
	row := ur.db.QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Email,
		u.Picture, u.PhoneNumber, u.Address,
		u.IsActive, u.IsAdmin, time.Now(), id,
	)

	u, err := ScanRowUsers(row)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
