package repository

import (
	"database/sql"
	"fmt"
	"nftgenie/backend/database"
	"nftgenie/backend/models"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// NFTRepository handles NFT database operations
type NFTRepository struct {
	db *sqlx.DB
}

// NewNFTRepository creates a new NFT repository
func NewNFTRepository() *NFTRepository {
	return &NFTRepository{
		db: database.DB,
	}
}

// Create creates a new NFT
func (r *NFTRepository) Create(nft *models.NFT) error {
	query := `
		INSERT INTO nfts (
			id, name, description, image_url, metadata_url,
			creator_id, owner_id, collection_id, contract_address,
			token_id, chain, transaction_hash, minted_at,
			views, likes, attributes, tags
		) VALUES (
			:id, :name, :description, :image_url, :metadata_url,
			:creator_id, :owner_id, :collection_id, :contract_address,
			:token_id, :chain, :transaction_hash, :minted_at,
			:views, :likes, :attributes, :tags
		)`
	
	_, err := r.db.NamedExec(query, nft)
	return err
}

// GetByID retrieves an NFT by ID
func (r *NFTRepository) GetByID(id uuid.UUID) (*models.NFT, error) {
	var nft models.NFT
	query := `
		SELECT n.*, 
			   u1.wallet_address as "creator.wallet_address",
			   u1.username as "creator.username",
			   u2.wallet_address as "owner.wallet_address",
			   u2.username as "owner.username"
		FROM nfts n
		LEFT JOIN users u1 ON n.creator_id = u1.id
		LEFT JOIN users u2 ON n.owner_id = u2.id
		WHERE n.id = $1`
	
	err := r.db.Get(&nft, query, id)
	if err != nil {
		return nil, err
	}
	
	// Increment view count
	r.IncrementViews(id)
	
	return &nft, nil
}

// GetAll retrieves all NFTs with pagination
func (r *NFTRepository) GetAll(limit, offset int) ([]*models.NFT, int, error) {
	var nfts []*models.NFT
	query := `
		SELECT n.*, 
			   u1.wallet_address as "creator.wallet_address",
			   u1.username as "creator.username",
			   u2.wallet_address as "owner.wallet_address",
			   u2.username as "owner.username"
		FROM nfts n
		LEFT JOIN users u1 ON n.creator_id = u1.id
		LEFT JOIN users u2 ON n.owner_id = u2.id
		ORDER BY n.created_at DESC
		LIMIT $1 OFFSET $2`
	
	err := r.db.Select(&nfts, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM nfts`
	err = r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, err
	}
	
	return nfts, total, nil
}

// GetByOwner retrieves NFTs by owner
func (r *NFTRepository) GetByOwner(ownerID uuid.UUID) ([]*models.NFT, error) {
	var nfts []*models.NFT
	query := `
		SELECT * FROM nfts 
		WHERE owner_id = $1 
		ORDER BY created_at DESC`
	
	err := r.db.Select(&nfts, query, ownerID)
	return nfts, err
}

// GetByCreator retrieves NFTs by creator
func (r *NFTRepository) GetByCreator(creatorID uuid.UUID) ([]*models.NFT, error) {
	var nfts []*models.NFT
	query := `
		SELECT * FROM nfts 
		WHERE creator_id = $1 
		ORDER BY created_at DESC`
	
	err := r.db.Select(&nfts, query, creatorID)
	return nfts, err
}

// GetByCollection retrieves NFTs by collection
func (r *NFTRepository) GetByCollection(collectionID uuid.UUID) ([]*models.NFT, error) {
	var nfts []*models.NFT
	query := `
		SELECT * FROM nfts 
		WHERE collection_id = $1 
		ORDER BY created_at DESC`
	
	err := r.db.Select(&nfts, query, collectionID)
	return nfts, err
}

// GetByTags retrieves NFTs by tags
func (r *NFTRepository) GetByTags(tags []string) ([]*models.NFT, error) {
	var nfts []*models.NFT
	query := `
		SELECT * FROM nfts 
		WHERE tags && $1 
		ORDER BY created_at DESC`
	
	err := r.db.Select(&nfts, query, pq.Array(tags))
	return nfts, err
}

// GetTrending retrieves trending NFTs
func (r *NFTRepository) GetTrending(limit int) ([]*models.NFT, error) {
	var nfts []*models.NFT
	query := `
		SELECT n.*, 
			   COUNT(DISTINCT ui.id) as interaction_count,
			   COUNT(DISTINCT ml.id) as listing_count
		FROM nfts n
		LEFT JOIN user_interactions ui ON n.id = ui.nft_id 
			AND ui.created_at > NOW() - INTERVAL '7 days'
		LEFT JOIN marketplace_listings ml ON n.id = ml.nft_id 
			AND ml.status = 'active'
		GROUP BY n.id
		ORDER BY (n.views * 0.3 + n.likes * 0.5 + 
				  COUNT(DISTINCT ui.id) * 0.2) DESC
		LIMIT $1`
	
	err := r.db.Select(&nfts, query, limit)
	return nfts, err
}

// Update updates an NFT
func (r *NFTRepository) Update(nft *models.NFT) error {
	query := `
		UPDATE nfts SET
			name = :name,
			description = :description,
			image_url = :image_url,
			metadata_url = :metadata_url,
			owner_id = :owner_id,
			attributes = :attributes,
			tags = :tags,
			updated_at = NOW()
		WHERE id = :id`
	
	_, err := r.db.NamedExec(query, nft)
	return err
}

// UpdateOwner updates NFT owner
func (r *NFTRepository) UpdateOwner(nftID, newOwnerID uuid.UUID) error {
	query := `
		UPDATE nfts SET
			owner_id = $1,
			updated_at = NOW()
		WHERE id = $2`
	
	_, err := r.db.Exec(query, newOwnerID, nftID)
	return err
}

// IncrementViews increments the view count
func (r *NFTRepository) IncrementViews(nftID uuid.UUID) error {
	query := `UPDATE nfts SET views = views + 1 WHERE id = $1`
	_, err := r.db.Exec(query, nftID)
	return err
}

// IncrementLikes increments the like count
func (r *NFTRepository) IncrementLikes(nftID uuid.UUID) error {
	query := `UPDATE nfts SET likes = likes + 1 WHERE id = $1`
	_, err := r.db.Exec(query, nftID)
	return err
}

// Delete deletes an NFT
func (r *NFTRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM nfts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Search searches NFTs by name or description
func (r *NFTRepository) Search(searchTerm string, limit, offset int) ([]*models.NFT, int, error) {
	var nfts []*models.NFT
	searchPattern := "%" + strings.ToLower(searchTerm) + "%"
	
	query := `
		SELECT * FROM nfts 
		WHERE LOWER(name) LIKE $1 
		   OR LOWER(description) LIKE $1
		   OR $2 = ANY(tags)
		ORDER BY views DESC, created_at DESC
		LIMIT $3 OFFSET $4`
	
	err := r.db.Select(&nfts, query, searchPattern, searchTerm, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	// Get total count
	var total int
	countQuery := `
		SELECT COUNT(*) FROM nfts 
		WHERE LOWER(name) LIKE $1 
		   OR LOWER(description) LIKE $1
		   OR $2 = ANY(tags)`
	err = r.db.Get(&total, countQuery, searchPattern, searchTerm)
	if err != nil {
		return nil, 0, err
	}
	
	return nfts, total, nil
}

// GetRecommendedForUser gets recommended NFTs for a user
func (r *NFTRepository) GetRecommendedForUser(userID uuid.UUID, limit int) ([]*models.NFT, error) {
	var nfts []*models.NFT
	
	// Complex recommendation query based on user interactions and preferences
	query := `
		WITH user_tags AS (
			SELECT UNNEST(n.tags) as tag, COUNT(*) as tag_count
			FROM user_interactions ui
			JOIN nfts n ON ui.nft_id = n.id
			WHERE ui.user_id = $1
			GROUP BY tag
			ORDER BY tag_count DESC
			LIMIT 10
		),
		user_creators AS (
			SELECT n.creator_id, COUNT(*) as creator_count
			FROM user_interactions ui
			JOIN nfts n ON ui.nft_id = n.id
			WHERE ui.user_id = $1
			GROUP BY n.creator_id
			ORDER BY creator_count DESC
			LIMIT 5
		)
		SELECT DISTINCT n.*,
			   (CASE WHEN n.creator_id IN (SELECT creator_id FROM user_creators) THEN 3 ELSE 0 END +
			    CASE WHEN EXISTS (SELECT 1 FROM user_tags ut WHERE ut.tag = ANY(n.tags)) THEN 2 ELSE 0 END +
			    (n.views * 0.0001) + (n.likes * 0.001)) as score
		FROM nfts n
		WHERE n.owner_id != $1
		  AND n.id NOT IN (
			  SELECT nft_id FROM user_interactions 
			  WHERE user_id = $1 AND interaction_type = 'purchase'
		  )
		ORDER BY score DESC, n.created_at DESC
		LIMIT $2`
	
	err := r.db.Select(&nfts, query, userID, limit)
	return nfts, err
}

// GetByContractAndToken retrieves an NFT by contract address and token ID
func (r *NFTRepository) GetByContractAndToken(contractAddress, tokenID string) (*models.NFT, error) {
	var nft models.NFT
	query := `
		SELECT * FROM nfts 
		WHERE contract_address = $1 AND token_id = $2`
	
	err := r.db.Get(&nft, query, contractAddress, tokenID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("NFT not found")
	}
	return &nft, err
}
