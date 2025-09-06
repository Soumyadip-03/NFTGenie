-- NFTGenie Database Schema
-- PostgreSQL Database Schema for NFT Marketplace

-- Create database (run this separately if needed)
-- CREATE DATABASE nftgenie;

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_address VARCHAR(42) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE,
    bio TEXT,
    profile_image VARCHAR(500),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_verified BOOLEAN DEFAULT FALSE,
    nonce VARCHAR(255) -- For signature verification
);

-- NFT Collections table
CREATE TABLE IF NOT EXISTS collections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    creator_id UUID REFERENCES users(id) ON DELETE CASCADE,
    contract_address VARCHAR(42),
    chain VARCHAR(50) DEFAULT 'polygonAmoy',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- NFTs table
CREATE TABLE IF NOT EXISTS nfts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(500) NOT NULL,
    metadata_url VARCHAR(500),
    creator_id UUID REFERENCES users(id) ON DELETE CASCADE,
    owner_id UUID REFERENCES users(id) ON DELETE CASCADE,
    collection_id UUID REFERENCES collections(id) ON DELETE CASCADE,
    contract_address VARCHAR(42),
    token_id VARCHAR(255),
    chain VARCHAR(50) DEFAULT 'polygonAmoy',
    transaction_hash VARCHAR(66),
    minted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    views INTEGER DEFAULT 0,
    likes INTEGER DEFAULT 0,
    attributes JSONB, -- Store NFT traits/attributes
    tags TEXT[] -- Array of tags for recommendations
);

-- Marketplace listings table
CREATE TABLE IF NOT EXISTS marketplace_listings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nft_id UUID REFERENCES nfts(id) ON DELETE CASCADE,
    seller_id UUID REFERENCES users(id) ON DELETE CASCADE,
    price DECIMAL(20, 8) NOT NULL,
    currency VARCHAR(10) DEFAULT 'MATIC',
    status VARCHAR(20) DEFAULT 'active', -- active, sold, cancelled
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sold_at TIMESTAMP,
    buyer_id UUID REFERENCES users(id) ON DELETE SET NULL
);

-- Transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type VARCHAR(20) NOT NULL, -- mint, transfer, purchase, list
    nft_id UUID REFERENCES nfts(id) ON DELETE CASCADE,
    from_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    to_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    transaction_hash VARCHAR(66),
    block_number BIGINT,
    price DECIMAL(20, 8),
    gas_fee DECIMAL(20, 8),
    status VARCHAR(20) DEFAULT 'pending', -- pending, confirmed, failed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User interactions table (for recommendations)
CREATE TABLE IF NOT EXISTS user_interactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    nft_id UUID REFERENCES nfts(id) ON DELETE CASCADE,
    interaction_type VARCHAR(20) NOT NULL, -- view, like, purchase, mint
    interaction_value FLOAT DEFAULT 1.0, -- Weight for recommendation system
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, nft_id, interaction_type)
);

-- User preferences table (for AI recommendations)
CREATE TABLE IF NOT EXISTS user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE UNIQUE,
    preferred_categories TEXT[],
    preferred_price_range JSONB, -- {"min": 0, "max": 100}
    preferred_creators UUID[],
    excluded_tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Recommendations cache table
CREATE TABLE IF NOT EXISTS recommendations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    nft_id UUID REFERENCES nfts(id) ON DELETE CASCADE,
    score FLOAT NOT NULL,
    reason TEXT,
    algorithm_version VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    UNIQUE(user_id, nft_id, algorithm_version)
);

-- Analytics table
CREATE TABLE IF NOT EXISTS analytics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    date DATE NOT NULL,
    total_users INTEGER DEFAULT 0,
    total_nfts INTEGER DEFAULT 0,
    total_transactions INTEGER DEFAULT 0,
    total_volume DECIMAL(20, 8) DEFAULT 0,
    active_users INTEGER DEFAULT 0,
    new_users INTEGER DEFAULT 0,
    trending_nfts JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(date)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_nfts_owner ON nfts(owner_id);
CREATE INDEX IF NOT EXISTS idx_nfts_creator ON nfts(creator_id);
CREATE INDEX IF NOT EXISTS idx_nfts_collection ON nfts(collection_id);
CREATE INDEX IF NOT EXISTS idx_nfts_tags ON nfts USING GIN(tags);
CREATE INDEX IF NOT EXISTS idx_nfts_attributes ON nfts USING GIN(attributes);
CREATE INDEX IF NOT EXISTS idx_marketplace_status ON marketplace_listings(status);
CREATE INDEX IF NOT EXISTS idx_marketplace_seller ON marketplace_listings(seller_id);
CREATE INDEX IF NOT EXISTS idx_transactions_user ON transactions(from_user_id, to_user_id);
CREATE INDEX IF NOT EXISTS idx_interactions_user ON user_interactions(user_id);
CREATE INDEX IF NOT EXISTS idx_interactions_nft ON user_interactions(nft_id);
CREATE INDEX IF NOT EXISTS idx_recommendations_user ON recommendations(user_id);
CREATE INDEX IF NOT EXISTS idx_users_wallet ON users(wallet_address);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply updated_at triggers (drop and recreate to avoid conflicts)
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_nfts_updated_at ON nfts;
CREATE TRIGGER update_nfts_updated_at BEFORE UPDATE ON nfts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_collections_updated_at ON collections;
CREATE TRIGGER update_collections_updated_at BEFORE UPDATE ON collections
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_marketplace_listings_updated_at ON marketplace_listings;
CREATE TRIGGER update_marketplace_listings_updated_at BEFORE UPDATE ON marketplace_listings
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_user_preferences_updated_at ON user_preferences;
CREATE TRIGGER update_user_preferences_updated_at BEFORE UPDATE ON user_preferences
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
