package database

import (
	"database/sql"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

// ScanRowUsers is a function to scan a row to a user.User
func ScanRowUsers(s Scanner) (models.User, error) {
	u := models.User{}
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
func ScanRowProducers(s Scanner) (models.Producer, error) {
	p := models.Producer{}

	var lastName, documentNumber, phoneNumber, birthDate, address sql.NullString

	err := s.Scan(
		&p.ID,
		&p.FirstName,
		&lastName,
		&documentNumber,
		&birthDate,
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
	p.DocumentNumber = documentNumber.String
	p.BirthDate = birthDate.String
	p.PhoneNumber = phoneNumber.String
	p.Address = address.String

	return p, nil
}

// ScanRowProduction is a function to scan a row to a production.Production
func ScanRowProduction(s Scanner) (models.Production, error) {
	pd := models.Production{}

	var latitude, longitude sql.NullFloat64
	var area, cultivatedArea sql.NullFloat64
	var loteNumber, entry, picture, cadastralRegistration, district sql.NullString

	err := s.Scan(
		&pd.ID,
		&pd.Producer,
		&loteNumber,
		&entry,
		&pd.Name,
		&pd.ProductionType,
		&area,
		&cultivatedArea,
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
	pd.Area = area.Float64
	pd.CultivatedArea = cultivatedArea.Float64
	pd.Latitude = latitude.Float64
	pd.Longitude = longitude.Float64

	return pd, nil
}

// ScanRowProduction is a function to scan a row to a production.ProductionResponse
func ScanRowProductionResponse(s Scanner) (models.ProductionResponse, error) {
	pdcs := models.ProductionResponse{}

	var latitude, longitude sql.NullFloat64
	var area, cultivatedArea sql.NullFloat64
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
		&area,
		&cultivatedArea,
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
	pdcs.Area = area.Float64
	pdcs.CultivatedArea = cultivatedArea.Float64
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
func ScanRowSection(s Scanner) (models.Section, error) {

	sct := models.Section{}

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

// ScanRowSectionResponse is a function to scan a row to a section.SectionResponse
/*
func ScanRowSectionResponse(s Scanner) (section.SectionResponse, error) {

	sct := section.SectionResponse{}

	var name sql.NullString

	var intakeID sql.NullInt64
	var intakeNumber, intakeName sql.NullString
	var intakeLatitude, intakeLongitude sql.NullFloat64

	err := s.Scan(
		&sct.ID,
		&sct.SectionNumber,
		&name,
		&intakeID,
		&intakeNumber,
		&intakeName,
		&intakeLatitude,
		&intakeLongitude,
		&sct.CreatedAt,
		&sct.UpdatedAt,
	)

	if err != nil {
		return sct, err
	}

	sct.Name = name.String
	sct.Intake.ID = intakeID.Int64
	sct.Intake.Number = intakeNumber.String
	sct.Intake.Name = intakeName.String
	sct.Intake.Latitude = intakeLatitude.Float64
	sct.Intake.Longitude = intakeLongitude.Float64

	return sct, nil
}
*/

// ScanRowIntake is a function to scan a row to a intake.Intake
func ScanRowIntake(s Scanner) (models.Intake, error) {

	ik := models.Intake{}

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
func ScanRowIntakeResponse(s Scanner) (models.IntakeResponse, error) {

	ik := models.IntakeResponse{}

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

// ScanRowIntakeProduction is a function to scan a row to a intake.IntakeProduction
func ScanRowIntakeProduction(s Scanner) (models.IntakeProduction, error) {

	ip := models.IntakeProduction{}

	var wateringOrder sql.NullInt64

	err := s.Scan(
		&ip.IntakeID,
		&ip.ProductionID,
		&wateringOrder,
	)

	if err != nil {
		return ip, err
	}

	ip.WateringOrder = wateringOrder.Int64

	return ip, nil
}

// ScanRowIntakeProductionResponse is a function to scan a row to a intake.IntakeProductionResponse
func ScanRowProductionIntakeResponse(s Scanner) (models.ProductionIntakeResponse, error) {
	pir := models.ProductionIntakeResponse{}

	var latitude, longitude sql.NullFloat64
	var area, cultivatedArea sql.NullFloat64
	var loteNumber, entry, picture, cadastralRegistration, district sql.NullString

	var producerId sql.NullInt64
	var producerFirstName, producerLastName sql.NullString
	var producerDocumentNumber, producerBirthDate sql.NullString
	var producerPhoneNumber, producerAddress sql.NullString

	var wateringOrder sql.NullInt64

	err := s.Scan(
		&pir.ID,
		&producerId,
		&producerFirstName,
		&producerLastName,
		&producerDocumentNumber,
		&producerBirthDate,
		&producerPhoneNumber,
		&producerAddress,
		&loteNumber,
		&entry,
		&pir.Name,
		&pir.ProductionType,
		&area,
		&cultivatedArea,
		&latitude,
		&longitude,
		&picture,
		&cadastralRegistration,
		&district,
		&wateringOrder,
		&pir.CreatedAt,
		&pir.UpdatedAt,
	)

	if err != nil {
		return pir, err
	}

	pir.LoteNumber = loteNumber.String
	pir.Entry = entry.String
	pir.Picture = picture.String
	pir.CadastralRegistration = cadastralRegistration.String
	pir.District = district.String
	pir.Area = area.Float64
	pir.CultivatedArea = cultivatedArea.Float64
	pir.WateringOrder = wateringOrder.Int64
	pir.Latitude = latitude.Float64
	pir.Longitude = longitude.Float64

	pir.Producer.ID = producerId.Int64
	pir.Producer.FirstName = producerFirstName.String
	pir.Producer.LastName = producerLastName.String
	pir.Producer.DocumentNumber = producerDocumentNumber.String
	pir.Producer.BirthDate = producerBirthDate.String
	pir.Producer.PhoneNumber = producerPhoneNumber.String
	pir.Producer.Address = producerAddress.String

	return pir, nil
}

// ScamRowTurn is a function to scan a row to a turn.Turn
func ScanRowTurn(s Scanner) (models.Turn, error) {

	t := models.Turn{}

	var turnHours sql.NullFloat64
	var endDate sql.NullString

	err := s.Scan(
		&t.ID,
		&t.StartDate,
		&turnHours,
		&endDate,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	if err != nil {
		return t, err
	}

	t.TurnHours = turnHours.Float64
	t.EndDate = endDate.String

	return t, nil
}

// ScanRowTurnResponse is a function to scan a row to a turn.TurnResponse
func ScanRowTurnResponse(s Scanner) (models.TurnResponse, error) {

	t := models.TurnResponse{}

	var turnHours sql.NullFloat64
	var endDate sql.NullString

	err := s.Scan(
		&t.ID,
		&t.StartDate,
		&turnHours,
		&endDate,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	if err != nil {
		return t, err
	}

	t.TurnHours = turnHours.Float64
	t.EndDate = endDate.String

	return t, nil
}

// ScanRowIntakeProductionResponse is a function to scan a row to a intake.IntakeProductionResponse
func ScanRowProductionTurnResponse(s Scanner) (models.ProductionTurnResponse, error) {
	ptr := models.ProductionTurnResponse{}

	var latitude, longitude sql.NullFloat64
	var area, cultivatedArea sql.NullFloat64
	var loteNumber, entry, picture, cadastralRegistration, district sql.NullString

	var producerId sql.NullInt64
	var producerFirstName, producerLastName sql.NullString
	var producerDocumentNumber, producerBirthDate sql.NullString
	var producerPhoneNumber, producerAddress sql.NullString

	var intakeId, wateringOrder sql.NullInt64
	var intakeNumber sql.NullString

	err := s.Scan(
		&ptr.ID,
		&producerId,
		&producerFirstName,
		&producerLastName,
		&producerDocumentNumber,
		&producerBirthDate,
		&producerPhoneNumber,
		&producerAddress,
		&loteNumber,
		&entry,
		&ptr.Name,
		&ptr.ProductionType,
		&area,
		&cultivatedArea,
		&latitude,
		&longitude,
		&picture,
		&cadastralRegistration,
		&district,
		&intakeId,
		&intakeNumber,
		&wateringOrder,
		&ptr.CreatedAt,
		&ptr.UpdatedAt,
	)

	if err != nil {
		return ptr, err
	}

	ptr.LoteNumber = loteNumber.String
	ptr.Entry = entry.String
	ptr.Picture = picture.String
	ptr.CadastralRegistration = cadastralRegistration.String
	ptr.District = district.String
	ptr.Area = area.Float64
	ptr.CultivatedArea = cultivatedArea.Float64

	ptr.IntakeID = intakeId.Int64
	ptr.IntakeNumber = intakeNumber.String
	ptr.WateringOrder = wateringOrder.Int64

	ptr.Latitude = latitude.Float64
	ptr.Longitude = longitude.Float64

	ptr.Producer.ID = producerId.Int64
	ptr.Producer.FirstName = producerFirstName.String
	ptr.Producer.LastName = producerLastName.String
	ptr.Producer.DocumentNumber = producerDocumentNumber.String
	ptr.Producer.BirthDate = producerBirthDate.String
	ptr.Producer.PhoneNumber = producerPhoneNumber.String
	ptr.Producer.Address = producerAddress.String

	return ptr, nil
}

// ScanRowTurnProduction is a function to scan a row to a turn.TurnProduction
func ScanRowTurnProduction(s Scanner) (models.TurnProduction, error) {

	tp := models.TurnProduction{}

	err := s.Scan(
		&tp.TurnID,
		&tp.ProductionID,
	)

	if err != nil {
		return tp, err
	}

	return tp, nil
}

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
