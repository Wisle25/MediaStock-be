CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    user_id UUID NOT NULL,
    total_amount BIGINT NOT NULL,
    purchased_at TIMESTAMP DEFAULT now() NOT NULL,
    
    -- Relation --
    CONSTRAINT fk_transactions_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transaction_items (
    id UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
    transaction_id UUID NOT NULL,
    asset_id UUID NOT NULL,
    purchased_at TIMESTAMP DEFAULT now() NOT NULL,
    
    -- Relations --
    CONSTRAINT fk_transaction_items_transaction FOREIGN KEY(transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
    CONSTRAINT fk_transaction_items_asset FOREIGN KEY(asset_id) REFERENCES assets(id) ON DELETE CASCADE
);
