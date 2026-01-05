-- Eliminar índices
DROP INDEX IF EXISTS idx_productions_area_coordinates;
DROP INDEX IF EXISTS idx_productions_cultivated_area_coordinates;

-- Restaurar columnas de punto único
ALTER TABLE productions
ADD COLUMN latitude DECIMAL(10,8) DEFAULT 0.0,
ADD COLUMN longitude DECIMAL(11,8) DEFAULT 0.0;

-- Eliminar columnas de polígonos
ALTER TABLE productions
DROP COLUMN IF EXISTS area_coordinates,
DROP COLUMN IF EXISTS cultivated_area_coordinates;
