CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    asset_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT fk_comments_asset FOREIGN KEY(asset_id) REFERENCES assets(id) ON DELETE CASCADE,
    CONSTRAINT fk_comments_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
