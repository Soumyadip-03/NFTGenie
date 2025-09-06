package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// User represents a user in the marketplace
type User struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	WalletAddress string     `db:"wallet_address" json:"wallet_address"`
	Username      *string    `db:"username" json:"username,omitempty"`
	Bio           *string    `db:"bio" json:"bio,omitempty"`
	ProfileImage  *string    `db:"profile_image" json:"profile_image,omitempty"`
	Email         *string    `db:"email" json:"email,omitempty"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
	IsVerified    bool       `db:"is_verified" json:"is_verified"`
	Nonce         *string    `db:"nonce" json:"-"`
}

// Collection represents an NFT collection
type Collection struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	Name            string     `db:"name" json:"name"`
	Description     *string    `db:"description" json:"description,omitempty"`
	CreatorID       uuid.UUID  `db:"creator_id" json:"creator_id"`
	ContractAddress *string    `db:"contract_address" json:"contract_address,omitempty"`
	Chain           string     `db:"chain" json:"chain"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
}

// NFT represents an NFT in the marketplace
type NFT struct {
	ID              uuid.UUID       `db:"id" json:"id"`
	Name            string          `db:"name" json:"name"`
	Description     *string         `db:"description" json:"description,omitempty"`
	ImageURL        string          `db:"image_url" json:"image_url"`
	MetadataURL     *string         `db:"metadata_url" json:"metadata_url,omitempty"`
	CreatorID       uuid.UUID       `db:"creator_id" json:"creator_id"`
	OwnerID         uuid.UUID       `db:"owner_id" json:"owner_id"`
	CollectionID    *uuid.UUID      `db:"collection_id" json:"collection_id,omitempty"`
	ContractAddress *string         `db:"contract_address" json:"contract_address,omitempty"`
	TokenID         *string         `db:"token_id" json:"token_id,omitempty"`
	Chain           string          `db:"chain" json:"chain"`
	TransactionHash *string         `db:"transaction_hash" json:"transaction_hash,omitempty"`
	MintedAt        *time.Time      `db:"minted_at" json:"minted_at,omitempty"`
	CreatedAt       time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at" json:"updated_at"`
	Views           int             `db:"views" json:"views"`
	Likes           int             `db:"likes" json:"likes"`
	Attributes      json.RawMessage `db:"attributes" json:"attributes,omitempty"`
	Tags            pq.StringArray  `db:"tags" json:"tags,omitempty"`
	
	// Joined fields (populated via joins)
	Creator         *User           `db:"-" json:"creator,omitempty"`
	Owner           *User           `db:"-" json:"owner,omitempty"`
	Collection      *Collection     `db:"-" json:"collection,omitempty"`
}

// MarketplaceListing represents an NFT listing in the marketplace
type MarketplaceListing struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	NFTID     uuid.UUID  `db:"nft_id" json:"nft_id"`
	SellerID  uuid.UUID  `db:"seller_id" json:"seller_id"`
	Price     float64    `db:"price" json:"price"`
	Currency  string     `db:"currency" json:"currency"`
	Status    string     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	SoldAt    *time.Time `db:"sold_at" json:"sold_at,omitempty"`
	BuyerID   *uuid.UUID `db:"buyer_id" json:"buyer_id,omitempty"`
	
	// Joined fields
	NFT       *NFT       `db:"-" json:"nft,omitempty"`
	Seller    *User      `db:"-" json:"seller,omitempty"`
	Buyer     *User      `db:"-" json:"buyer,omitempty"`
}

// Transaction represents a blockchain transaction
type Transaction struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	Type            string     `db:"type" json:"type"`
	NFTID           uuid.UUID  `db:"nft_id" json:"nft_id"`
	FromUserID      *uuid.UUID `db:"from_user_id" json:"from_user_id,omitempty"`
	ToUserID        *uuid.UUID `db:"to_user_id" json:"to_user_id,omitempty"`
	TransactionHash *string    `db:"transaction_hash" json:"transaction_hash,omitempty"`
	BlockNumber     *int64     `db:"block_number" json:"block_number,omitempty"`
	Price           *float64   `db:"price" json:"price,omitempty"`
	GasFee          *float64   `db:"gas_fee" json:"gas_fee,omitempty"`
	Status          string     `db:"status" json:"status"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
}

// UserInteraction represents user interactions with NFTs
type UserInteraction struct {
	ID               uuid.UUID `db:"id" json:"id"`
	UserID           uuid.UUID `db:"user_id" json:"user_id"`
	NFTID            uuid.UUID `db:"nft_id" json:"nft_id"`
	InteractionType  string    `db:"interaction_type" json:"interaction_type"`
	InteractionValue float64   `db:"interaction_value" json:"interaction_value"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

// UserPreferences represents user preferences for recommendations
type UserPreferences struct {
	ID                  uuid.UUID         `db:"id" json:"id"`
	UserID              uuid.UUID         `db:"user_id" json:"user_id"`
	PreferredCategories pq.StringArray    `db:"preferred_categories" json:"preferred_categories,omitempty"`
	PreferredPriceRange json.RawMessage   `db:"preferred_price_range" json:"preferred_price_range,omitempty"`
	PreferredCreators   pq.StringArray    `db:"preferred_creators" json:"preferred_creators,omitempty"`
	ExcludedTags        pq.StringArray    `db:"excluded_tags" json:"excluded_tags,omitempty"`
	CreatedAt           time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time         `db:"updated_at" json:"updated_at"`
}

// Recommendation represents a cached recommendation
type Recommendation struct {
	ID               uuid.UUID  `db:"id" json:"id"`
	UserID           uuid.UUID  `db:"user_id" json:"user_id"`
	NFTID            uuid.UUID  `db:"nft_id" json:"nft_id"`
	Score            float64    `db:"score" json:"score"`
	Reason           *string    `db:"reason" json:"reason,omitempty"`
	AlgorithmVersion string     `db:"algorithm_version" json:"algorithm_version"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	ExpiresAt        *time.Time `db:"expires_at" json:"expires_at,omitempty"`
	
	// Joined field
	NFT              *NFT       `db:"-" json:"nft,omitempty"`
}

// Analytics represents daily analytics
type Analytics struct {
	ID               uuid.UUID       `db:"id" json:"id"`
	Date             time.Time       `db:"date" json:"date"`
	TotalUsers       int             `db:"total_users" json:"total_users"`
	TotalNFTs        int             `db:"total_nfts" json:"total_nfts"`
	TotalTransactions int            `db:"total_transactions" json:"total_transactions"`
	TotalVolume      float64         `db:"total_volume" json:"total_volume"`
	ActiveUsers      int             `db:"active_users" json:"active_users"`
	NewUsers         int             `db:"new_users" json:"new_users"`
	TrendingNFTs     json.RawMessage `db:"trending_nfts" json:"trending_nfts,omitempty"`
	CreatedAt        time.Time       `db:"created_at" json:"created_at"`
}

// PriceRange represents a price range for preferences
type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// NFTAttribute represents an NFT attribute/trait
type NFTAttribute struct {
	TraitType string      `json:"trait_type"`
	AttrValue interface{} `json:"value"` // Renamed from Value to avoid conflict
}

// Helper methods

// NewUser creates a new user with defaults
func NewUser(walletAddress string) *User {
	return &User{
		ID:            uuid.New(),
		WalletAddress: walletAddress,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsVerified:    false,
	}
}

// NewNFT creates a new NFT with defaults
func NewNFT(name, imageURL string, creatorID, ownerID uuid.UUID) *NFT {
	return &NFT{
		ID:        uuid.New(),
		Name:      name,
		ImageURL:  imageURL,
		CreatorID: creatorID,
		OwnerID:   ownerID,
		Chain:     "polygonAmoy",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Views:     0,
		Likes:     0,
	}
}

// NewMarketplaceListing creates a new marketplace listing
func NewMarketplaceListing(nftID, sellerID uuid.UUID, price float64) *MarketplaceListing {
	return &MarketplaceListing{
		ID:        uuid.New(),
		NFTID:     nftID,
		SellerID:  sellerID,
		Price:     price,
		Currency:  "MATIC",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
