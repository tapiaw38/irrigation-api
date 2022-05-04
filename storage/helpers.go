package storage

import (
	"database/sql"
	"time"

	"github.com/tapiaw38/irrigation-api/models/producer"
	"github.com/tapiaw38/irrigation-api/models/user"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

// ScanRowUsers is a function to scan a row to a user.User
func ScanRowUsers(s Scanner) (user.User, error) {
	u := user.User{}
	var lastName, picture, phoneNumber, address sql.NullString

	err := s.Scan(
		&u.ID,
		&u.FirstName,
		&lastName,
		&u.Username,
		&u.Email,
		&picture,
		&phoneNumber,
		&address,
		&u.IsActive,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	u.LastName = lastName.String
	u.Picture = picture.String
	u.PhoneNumber = phoneNumber.String
	u.Address = address.String

	return u, nil
}

// ScanRowProducers is a function to scan a row to a producer.Producer
func ScanRowProducers(s Scanner) (producer.Producer, error) {
	p := producer.Producer{}

	var lastName, phoneNumber, address sql.NullString

	err := s.Scan(
		&p.ID,
		&p.FirstName,
		&lastName,
		&p.DocumentNumber,
		&p.BirthDate,
		&p.Address,
		&phoneNumber,
		&address,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return p, err
	}

	p.LastName = lastName.String
	p.PhoneNumber = phoneNumber.String
	p.Address = address.String

	return p, nil
}

// helper function to control the null fields

// StringToNull is a function to convert a string to a sql.NullString
func StringToNull(str string) sql.NullString {
	null := sql.NullString{String: str}
	if null.String == "" {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}

// IntToNull is a function to convert a int to a sql.NullInt64
func IntToNull(i int64) sql.NullInt64 {
	null := sql.NullInt64{Int64: i}
	if null.Int64 == 0 {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}

// FloatToNull is a function to convert a float to a sql.NullFloat64
func FloatToNull(f float64) sql.NullFloat64 {
	null := sql.NullFloat64{Float64: f}
	if null.Float64 == 0 {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}

func BoolToNull(b bool) sql.NullBool {
	null := sql.NullBool{Bool: b}
	if !null.Bool {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}

// ParsinTime is a function to convert a time.Time to a sql.NullTime
func ParsingTime(t time.Time) string {
	return t.Format("RFC3339")
}
