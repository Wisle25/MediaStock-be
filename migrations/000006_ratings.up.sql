CREATE TABLE IF NOT EXISTS ratings (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),   
    asset_id UUID NOT NULL,
    user_id UUID NOT NULL,
    score SMALLINT NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT fk_ratings_asset FOREIGN KEY(asset_id) REFERENCES assets(id) ON DELETE CASCADE,
    CONSTRAINT fk_ratings_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
