CREATE TABLE IF NOT EXISTS favorites (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    asset_id UUID NOT NULL,
    user_id UUID NOT NULL,
    
    CONSTRAINT fk_favorite_asset FOREIGN KEY(asset_id) REFERENCES assets(id) ON DELETE CASCADE,
    CONSTRAINT fk_favorite_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
