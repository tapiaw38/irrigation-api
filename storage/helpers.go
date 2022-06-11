package storage

import (
	"database/sql"
	"time"

	"github.com/tapiaw38/irrigation-api/models/intake"
	"github.com/tapiaw38/irrigation-api/models/producer"
	"github.com/tapiaw38/irrigation-api/models/production"
	"github.com/tapiaw38/irrigation-api/models/section"
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
		&phoneNumber,
		&address,
		&p.IsActive,
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

// ScanRowProduction is a function to scan a row to a production.Production
func ScanRowProduction(s Scanner) (production.Production, error) {
	pd := production.Production{}

	var latitude, longitude sql.NullFloat64
	var loteNumber, entry, picture, cadastralRegistration, district sql.NullString

	err := s.Scan(
		&pd.ID,
		&pd.Producer,
		&loteNumber,
		&entry,
		&pd.Name,
		&pd.ProductionType,
		&pd.Area,
		&latitude,
		&longitude,
		&picture,
		&cadastralRegistration,
		&district,
		&pd.CreatedAt,
		&pd.UpdatedAt,
	)

	if err != nil {
		return pd, err
	}

	pd.LoteNumber = loteNumber.String
	pd.Entry = entry.String
	pd.Picture = picture.String
	pd.CadastralRegistration = cadastralRegistration.String
	pd.District = district.String
	pd.Latitude = latitude.Float64
	pd.Longitude = longitude.Float64

	return pd, nil
}

// ScanRowProduction is a function to scan a row to a production.ProductionResponse
func ScanRowProductionResponse(s Scanner) (production.ProductionResponse, error) {
	pdcs := production.ProductionResponse{}

	var latitude, longitude sql.NullFloat64
	var loteNumber, entry, picture, cadastralRegistration, district sql.NullString

	var producerId sql.NullInt64
	var producerFirstName, producerLastName sql.NullString
	var producerDocumentNumber, producerBirthDate sql.NullString
	var producerPhoneNumber, producerAddress sql.NullString

	err := s.Scan(
		&pdcs.ID,
		&producerId,
		&producerFirstName,
		&producerLastName,
		&producerDocumentNumber,
		&producerBirthDate,
		&producerPhoneNumber,
		&producerAddress,
		&loteNumber,
		&entry,
		&pdcs.Name,
		&pdcs.ProductionType,
		&pdcs.Area,
		&latitude,
		&longitude,
		&picture,
		&cadastralRegistration,
		&district,
		&pdcs.CreatedAt,
		&pdcs.UpdatedAt,
	)

	if err != nil {
		return pdcs, err
	}

	pdcs.LoteNumber = loteNumber.String
	pdcs.Entry = entry.String
	pdcs.Picture = picture.String
	pdcs.CadastralRegistration = cadastralRegistration.String
	pdcs.District = district.String
	pdcs.Latitude = latitude.Float64
	pdcs.Longitude = longitude.Float64

	pdcs.Producer.ID = producerId.Int64
	pdcs.Producer.FirstName = producerFirstName.String
	pdcs.Producer.LastName = producerLastName.String
	pdcs.Producer.DocumentNumber = producerDocumentNumber.String
	pdcs.Producer.BirthDate = producerBirthDate.String
	pdcs.Producer.PhoneNumber = producerPhoneNumber.String
	pdcs.Producer.Address = producerAddress.String

	return pdcs, nil
}

// ScanRowSection is a function to scan a row to a section.Section
func ScanRowSection(s Scanner) (section.Section, error) {

	sct := section.Section{}

	var name sql.NullString

	err := s.Scan(
		&sct.ID,
		&sct.SectionNumber,
		&name,
		&sct.CreatedAt,
		&sct.UpdatedAt,
	)

	if err != nil {
		return sct, err
	}

	sct.Name = name.String

	return sct, nil
}

// ScanRowIntake is a function to scan a row to a intake.Intake
func ScanRowIntake(s Scanner) (intake.Intake, error) {

	ik := intake.Intake{}

	var latitude, longitude sql.NullFloat64
	var name sql.NullString

	err := s.Scan(
		&ik.ID,
		&ik.IntakeNumber,
		&name,
		&ik.Section,
		&latitude,
		&longitude,
		&ik.CreatedAt,
		&ik.UpdatedAt,
	)

	if err != nil {
		return ik, err
	}

	ik.Name = name.String
	ik.Latitude = latitude.Float64
	ik.Longitude = longitude.Float64

	return ik, nil
}

// ScanRowIntakeResponse is a function to scan a row to a intake.IntakeResponse
func ScanRowIntakeResponse(s Scanner) (intake.IntakeResponse, error) {

	ik := intake.IntakeResponse{}

	var latitude, longitude sql.NullFloat64
	var name sql.NullString

	var sectionId sql.NullInt64
	var sectionNumber, sectionName sql.NullString

	err := s.Scan(
		&ik.ID,
		&sectionId,
		&sectionNumber,
		&sectionName,
		&ik.IntakeNumber,
		&name,
		&latitude,
		&longitude,
		&ik.CreatedAt,
		&ik.UpdatedAt,
	)

	if err != nil {
		return ik, err
	}

	ik.Name = name.String
	ik.Latitude = latitude.Float64
	ik.Longitude = longitude.Float64

	ik.Section.ID = sectionId.Int64
	ik.Section.SectionNumber = sectionNumber.String
	ik.Section.Name = sectionName.String

	return ik, nil
}

/*** helper function to control the null fields ***/

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
