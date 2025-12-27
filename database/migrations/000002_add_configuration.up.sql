-- Create configuration table
CREATE TABLE IF NOT EXISTS configurations (
    id BIGSERIAL PRIMARY KEY,
    default_latitude DOUBLE PRECISION NOT NULL DEFAULT -28.065752,
    default_longitude DOUBLE PRECISION NOT NULL DEFAULT -67.564368,
    default_zoom INTEGER NOT NULL DEFAULT 13,
    watering_hour_factor DOUBLE PRECISION NOT NULL DEFAULT 2.0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Insert default configuration
INSERT INTO configurations (default_latitude, default_longitude, default_zoom, watering_hour_factor)
VALUES (-28.065752, -67.564368, 13, 2.0);
