-- Migration script: 000001_init.down.sql

-- Drop the url_clicks table
DROP TABLE IF EXISTS url_clicks;

-- Drop the unique index on the short_code column of the urls table
DROP INDEX IF EXISTS idx_short_code;

-- Drop the urls table
DROP TABLE IF EXISTS urls;

-- Drop the sessions table
DROP TABLE IF EXISTS sessions;

-- Drop the unique index on the email column of the users table
DROP INDEX IF EXISTS idx_users_email;

-- Drop the users table
DROP TABLE IF EXISTS users;

-- Disable the uuid-ossp extension
DROP EXTENSION IF EXISTS "uuid-ossp";