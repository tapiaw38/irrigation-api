package utils

import (
	"encoding/json"
	"math"

	"github.com/tapiaw38/irrigation-api/models"
)

const earthRadiusKm = 6371.0

func CalculatePolygonArea(coordinatesJSON string) (float64, error) {
	if coordinatesJSON == "" {
		return 0, nil
	}

	var coordinates []models.Coordinate
	err := json.Unmarshal([]byte(coordinatesJSON), &coordinates)
	if err != nil {
		return 0, err
	}

	if len(coordinates) < 3 {
		return 0, nil
	}

	areaM2 := calculateSphericalPolygonArea(coordinates)
	areaHa := areaM2 / 10000.0

	return areaHa, nil
}

func calculateSphericalPolygonArea(coordinates []models.Coordinate) float64 {
	n := len(coordinates)
	if n < 3 {
		return 0
	}

	if coordinates[0].Lat != coordinates[n-1].Lat || coordinates[0].Lng != coordinates[n-1].Lng {
		coordinates = append(coordinates, coordinates[0])
		n++
	}

	var area float64

	var centerLat, centerLng float64
	for _, coord := range coordinates[:n-1] {
		centerLat += coord.Lat
		centerLng += coord.Lng
	}
	centerLat /= float64(n - 1)
	centerLng /= float64(n - 1)

	centerLatRad := degreesToRadians(centerLat)

	metersPerDegreeLat := 111132.92 - 559.82*math.Cos(2*centerLatRad) + 1.175*math.Cos(4*centerLatRad)
	metersPerDegreeLng := 111412.84*math.Cos(centerLatRad) - 93.5*math.Cos(3*centerLatRad)

	for i := 0; i < n-1; i++ {
		x1 := coordinates[i].Lng * metersPerDegreeLng
		y1 := coordinates[i].Lat * metersPerDegreeLat
		x2 := coordinates[i+1].Lng * metersPerDegreeLng
		y2 := coordinates[i+1].Lat * metersPerDegreeLat

		area += x1*y2 - x2*y1
	}

	area = math.Abs(area) / 2.0

	return area
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func radiansToDegrees(radians float64) float64 {
	return radians * 180.0 / math.Pi
}

func ValidatePolygon(coordinatesJSON string) bool {
	if coordinatesJSON == "" {
		return false
	}

	var coordinates []models.Coordinate
	err := json.Unmarshal([]byte(coordinatesJSON), &coordinates)
	if err != nil {
		return false
	}

	return len(coordinates) >= 3
}

func ValidateCultivatedAreaWithinTotal(cultivatedArea, totalArea float64) bool {
	return cultivatedArea <= totalArea
}
