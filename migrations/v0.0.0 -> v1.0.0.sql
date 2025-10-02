-- Example initial schema migration: empty -> v1.0.0
-- Creates a PEOPLE table with a few basic attributes.

CREATE TABLE IF NOT EXISTS PEOPLE (
  NAME VARCHAR(255) NOT NULL,
  TITLE VARCHAR(255),
  EMAIL_ADDRESS VARCHAR(320) NOT NULL UNIQUE
);
