-- Migration: v1.0.0 -> v1.1.0
-- Adds an auto-incrementing unique ID column to PEOPLE.

-- Note: Uses Postgres BIGSERIAL which creates a sequence and sets
-- a default nextval(...) for the column.

ALTER TABLE IF EXISTS PEOPLE
  ADD COLUMN ID BIGSERIAL PRIMARY KEY;
