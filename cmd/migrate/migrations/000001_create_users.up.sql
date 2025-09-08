CREATE EXTENSION IF NOT EXISTS citext;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password bytea NOT NULL,
    email citext UNIQUE NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

