CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_users_email ON users(email);

CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE urls (
    id SERIAL PRIMARY KEY,     
    user_id UUID REFERENCES users(id),       
    original_url TEXT NOT NULL,       
    short_code VARCHAR(5) NOT NULL, 
    click_count INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL 
);

CREATE UNIQUE INDEX idx_short_code ON urls(short_code);

CREATE TABLE url_clicks (
    id SERIAL PRIMARY KEY,        
    url_id INT REFERENCES urls(id),
    ip_address VARCHAR(45),        
    user_agent TEXT,               
    referer TEXT,                  
    clicked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);