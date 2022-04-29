package storage

import (
	"github.com/tapiaw38/irrigation-api/models/user"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

func ScanRowUsers(s Scanner) (user.User, error) {
	u := user.User{}

	err := s.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Picture,
		&u.Address,
		&u.IsActive,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}
