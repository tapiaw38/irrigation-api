package database

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

// ScanRowConfiguration scans a row from the configurations table
func ScanRowConfiguration(s Scanner) (models.Configuration, error) {
	var config models.Configuration

	err := s.Scan(
		&config.ID,
		&config.DefaultLocation.Latitude,
		&config.DefaultLocation.Longitude,
		&config.DefaultLocation.Zoom,
		&config.WateringHourFactor,
		&config.CreatedAt,
		&config.UpdatedAt,
	)

	if err != nil {
		return models.Configuration{}, err
	}

	return config, nil
}

// GetConfiguration returns the system configuration
func (ss *PostgresRepository) GetConfiguration(ctx context.Context) (models.Configuration, error) {
	q := `
	SELECT id, default_latitude, default_longitude, default_zoom,
	       watering_hour_factor, created_at, updated_at
		FROM configurations
		LIMIT 1;
	`

	row := ss.db.QueryRowContext(ctx, q)

	config, err := ScanRowConfiguration(row)

	if err != nil {
		log.Println(err)
		return models.Configuration{}, err
	}

	return config, nil
}

// UpdateConfiguration updates the system configuration
func (ss *PostgresRepository) UpdateConfiguration(ctx context.Context, config *models.Configuration) (models.Configuration, error) {
	q := `
	UPDATE configurations
		SET default_latitude = $1,
		    default_longitude = $2,
		    default_zoom = $3,
		    watering_hour_factor = $4,
		    updated_at = $5
		WHERE id = $6
		RETURNING id, default_latitude, default_longitude, default_zoom,
		          watering_hour_factor, created_at, updated_at;
	`

	row := ss.db.QueryRowContext(
		ctx, q,
		config.DefaultLocation.Latitude,
		config.DefaultLocation.Longitude,
		config.DefaultLocation.Zoom,
		config.WateringHourFactor,
		time.Now(),
		config.ID,
	)

	updatedConfig, err := ScanRowConfiguration(row)

	if err != nil {
		log.Println(err)
		return models.Configuration{}, err
	}

	return updatedConfig, nil
}
