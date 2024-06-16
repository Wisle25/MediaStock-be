CREATE TABLE IF NOT EXISTS assets(
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    owner_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    file_path TEXT NOT NULL,
    file_watermark_path TEXT NOT NULL,
    description VARCHAR(255) NOT NULL,
    details TEXT,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL,
    
    -- Relation --
    CONSTRAINT fk_assets_user_owner FOREIGN KEY(owner_id) REFERENCES users(id) ON DELETE CASCADE
);
