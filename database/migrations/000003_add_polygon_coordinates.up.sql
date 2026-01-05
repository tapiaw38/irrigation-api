-- Agregar columnas para almacenar coordenadas de polígonos
ALTER TABLE productions
ADD COLUMN area_coordinates JSONB,
ADD COLUMN cultivated_area_coordinates JSONB;

-- Eliminar las columnas de punto único (latitude, longitude)
ALTER TABLE productions
DROP COLUMN IF EXISTS latitude,
DROP COLUMN IF EXISTS longitude;

-- Crear índices para mejorar performance en queries JSONB
CREATE INDEX idx_productions_area_coordinates ON productions USING GIN (area_coordinates);
CREATE INDEX idx_productions_cultivated_area_coordinates ON productions USING GIN (cultivated_area_coordinates);
