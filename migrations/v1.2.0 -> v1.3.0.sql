-- Migration: v1.2.0 -> v1.3.0
-- Adds an index on PLACES(NAME) to speed up lookups by place name.

-- Create index on PLACES.NAME
CREATE INDEX IF NOT EXISTS idx_places_name ON PLACES (NAME);
