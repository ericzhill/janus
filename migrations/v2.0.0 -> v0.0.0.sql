-- Migration: v2.0.0 -> empty
-- Drops all application tables to return to an empty schema.
-- Drop order respects foreign key dependencies.

-- 1) Drop join table first (depends on PEOPLE and PLACES)
DROP TABLE PERSON_PLACES;

-- 2) Drop PLACES (has an index that will be dropped automatically)
DROP TABLE PLACES;

-- 3) Drop PEOPLE
DROP TABLE PEOPLE;
