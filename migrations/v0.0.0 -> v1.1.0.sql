-- Example initial schema migration: empty -> v1.1.0
-- Creates a PEOPLE table with an auto-incrementing ID and a few basic attributes.

CREATE TABLE IF NOT EXISTS PEOPLE (
  ID BIGSERIAL PRIMARY KEY,
  NAME VARCHAR(255) NOT NULL,
  TITLE VARCHAR(255),
  EMAIL_ADDRESS VARCHAR(320) NOT NULL UNIQUE
);
