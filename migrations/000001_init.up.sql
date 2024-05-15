CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Board table
CREATE TABLE boards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE access_level_enum AS ENUM('viewer', 'editor', 'admin');

-- Permission table
CREATE TABLE board_permissions (
    board_id UUID NOT NULL REFERENCES boards (id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    access_level access_level_enum NOT NULL,
    PRIMARY KEY (board_id, user_id)
);

CREATE TYPE status_card_enum AS ENUM('to do', 'doing', 'done');

-- Cards of column table
CREATE TABLE cards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP WITH TIME ZONE,
    status_card status_card_enum NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Board cards table
CREATE TABLE board_cards (
    board_id UUID NOT NULL REFERENCES boards (id) ON DELETE CASCADE,
    card_id UUID NOT NULL REFERENCES cards (id) ON DELETE CASCADE,
    PRIMARY KEY (board_id, card_id)
);

CREATE TABLE confirmation_codes (
    email TEXT,
    code TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes to improve query performance
CREATE INDEX idx_user_email ON users (email);