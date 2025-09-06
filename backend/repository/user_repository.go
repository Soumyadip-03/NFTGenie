package repository

import (
	"database/sql"
	"fmt"
	"nftgenie/backend/database"
	"nftgenie/backend/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserRepository handles user database operations
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (
			id, wallet_address, username, bio, profile_image,
			email, is_verified, nonce
		) VALUES (
			:id, :wallet_address, :username, :bio, :profile_image,
			:email, :is_verified, :nonce
		)`
	
	_, err := r.db.NamedExec(query, user)
	return err
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	
	err := r.db.Get(&user, query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return &user, err
}

// GetByWalletAddress retrieves a user by wallet address
func (r *UserRepository) GetByWalletAddress(walletAddress string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE wallet_address = $1`
	
	err := r.db.Get(&user, query, walletAddress)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist yet
	}
	return &user, err
}

// GetByUsername retrieves a user by username
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE username = $1`
	
	err := r.db.Get(&user, query, username)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return &user, err
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users SET
			username = :username,
			bio = :bio,
			profile_image = :profile_image,
			email = :email,
			is_verified = :is_verified,
			updated_at = NOW()
		WHERE id = :id`
	
	_, err := r.db.NamedExec(query, user)
	return err
}

// UpdateNonce updates user's nonce for signature verification
func (r *UserRepository) UpdateNonce(userID uuid.UUID, nonce string) error {
	query := `UPDATE users SET nonce = $1 WHERE id = $2`
	_, err := r.db.Exec(query, nonce, userID)
	return err
}

// Delete deletes a user
func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// GetAll retrieves all users with pagination
func (r *UserRepository) GetAll(limit, offset int) ([]*models.User, int, error) {
	var users []*models.User
	query := `
		SELECT * FROM users 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	
	err := r.db.Select(&users, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM users`
	err = r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, err
	}
	
	return users, total, nil
}

// GetTopCreators retrieves top creators by NFT count
func (r *UserRepository) GetTopCreators(limit int) ([]*models.User, error) {
	var users []*models.User
	query := `
		SELECT u.*, COUNT(n.id) as nft_count
		FROM users u
		LEFT JOIN nfts n ON u.id = n.creator_id
		GROUP BY u.id
		HAVING COUNT(n.id) > 0
		ORDER BY nft_count DESC
		LIMIT $1`
	
	err := r.db.Select(&users, query, limit)
	return users, err
}

// GetTopCollectors retrieves top collectors by NFT ownership
func (r *UserRepository) GetTopCollectors(limit int) ([]*models.User, error) {
	var users []*models.User
	query := `
		SELECT u.*, COUNT(n.id) as nft_count, SUM(n.likes) as total_likes
		FROM users u
		LEFT JOIN nfts n ON u.id = n.owner_id
		GROUP BY u.id
		HAVING COUNT(n.id) > 0
		ORDER BY nft_count DESC, total_likes DESC
		LIMIT $1`
	
	err := r.db.Select(&users, query, limit)
	return users, err
}

// UsernameExists checks if a username already exists
func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := r.db.Get(&exists, query, username)
	return exists, err
}

// CreateOrUpdate creates a new user or updates existing one
func (r *UserRepository) CreateOrUpdate(user *models.User) error {
	query := `
		INSERT INTO users (
			id, wallet_address, username, bio, profile_image,
			email, is_verified, nonce
		) VALUES (
			:id, :wallet_address, :username, :bio, :profile_image,
			:email, :is_verified, :nonce
		)
		ON CONFLICT (wallet_address) 
		DO UPDATE SET
			username = EXCLUDED.username,
			bio = EXCLUDED.bio,
			profile_image = EXCLUDED.profile_image,
			email = EXCLUDED.email,
			updated_at = NOW()
		RETURNING id`
	
	var userID uuid.UUID
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	if rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			return err
		}
		user.ID = userID
	}
	
	return nil
}

// GetUserStats retrieves user statistics
func (r *UserRepository) GetUserStats(userID uuid.UUID) (map[string]interface{}, error) {
	stats := make(map[string]interface{})
	
	// Get NFT counts
	var nftStats struct {
		Created int `db:"created"`
		Owned   int `db:"owned"`
	}
	
	query := `
		SELECT 
			(SELECT COUNT(*) FROM nfts WHERE creator_id = $1) as created,
			(SELECT COUNT(*) FROM nfts WHERE owner_id = $1) as owned`
	
	err := r.db.Get(&nftStats, query, userID)
	if err != nil {
		return nil, err
	}
	
	stats["nfts_created"] = nftStats.Created
	stats["nfts_owned"] = nftStats.Owned
	
	// Get transaction stats
	var txStats struct {
		TotalSales     int     `db:"total_sales"`
		TotalPurchases int     `db:"total_purchases"`
		TotalVolume    float64 `db:"total_volume"`
	}
	
	query = `
		SELECT 
			(SELECT COUNT(*) FROM transactions WHERE from_user_id = $1 AND type = 'transfer') as total_sales,
			(SELECT COUNT(*) FROM transactions WHERE to_user_id = $1 AND type = 'transfer') as total_purchases,
			(SELECT COALESCE(SUM(price), 0) FROM transactions WHERE from_user_id = $1 OR to_user_id = $1) as total_volume`
	
	err = r.db.Get(&txStats, query, userID)
	if err != nil {
		return nil, err
	}
	
	stats["total_sales"] = txStats.TotalSales
	stats["total_purchases"] = txStats.TotalPurchases
	stats["total_volume"] = txStats.TotalVolume
	
	// Get interaction stats
	var interactionCount int
	query = `SELECT COUNT(*) FROM user_interactions WHERE user_id = $1`
	err = r.db.Get(&interactionCount, query, userID)
	if err != nil {
		return nil, err
	}
	
	stats["total_interactions"] = interactionCount
	
	return stats, nil
}
