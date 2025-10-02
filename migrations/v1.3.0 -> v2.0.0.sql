-- Migration: v1.3.0 -> v2.0.0
-- Replaces PEOPLE.NAME with FIRST_NAME and LAST_NAME. Attempts to split existing
-- NAME values on the first space to backfill FIRST_NAME and LAST_NAME.
-- Notes:
--  - If there is no space in NAME, FIRST_NAME will get the full NAME and
--    LAST_NAME will be NULL.
--  - If there are multiple spaces, FIRST_NAME is the first token and LAST_NAME
--    is everything after the first space.

-- 1) Add new columns FIRST_NAME and LAST_NAME
ALTER PEOPLE
  ADD COLUMN FIRST_NAME VARCHAR(255),
  ADD COLUMN LAST_NAME VARCHAR(255);

-- 2) Backfill data from NAME by splitting on the first space
UPDATE PEOPLE
SET
  FIRST_NAME = CASE
                 WHEN NAME IS NULL OR btrim(NAME) = '' THEN NULL
                 ELSE split_part(btrim(NAME), ' ', 1)
               END,
  LAST_NAME = CASE
                WHEN NAME IS NULL OR btrim(NAME) = '' THEN NULL
                WHEN position(' ' in btrim(NAME)) = 0 THEN NULL
                ELSE btrim(regexp_replace(btrim(NAME), '^\S+\s+', ''))
              END;

-- 3) Drop the old NAME column
ALTER PEOPLE
  DROP COLUMN IF EXISTS NAME;
