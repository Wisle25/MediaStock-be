CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    avatar_link TEXT DEFAULT '',
    is_verified BOOLEAN DEFAULT FALSE,
    role VARCHAR(20) DEFAULT 'Special Guest' CHECK (role IN ('Special Guest', 'Seller', 'Admin'))
);

CREATE UNIQUE INDEX users_email_idx ON users(email);
CREATE UNIQUE INDEX users_username_idx ON users(username);
