package main

import (
	"fmt"
	"log"
	"time"

	"gofr.dev/pkg/gofr"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"nftgenie/backend/database"
	"nftgenie/backend/models"
	"nftgenie/backend/repository"
	"nftgenie/backend/services"
)

// NFT represents an NFT in our marketplace
type NFT struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Creator     string    `json:"creator"`
	Owner       string    `json:"owner"`
	Price       string    `json:"price"`
	Chain       string    `json:"chain"`
	ContractAddress string `json:"contract_address"`
	TokenID     string    `json:"token_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// User represents a user in our marketplace
type User struct {
	ID            string    `json:"id"`
	WalletAddress string    `json:"wallet_address"`
	Username      string    `json:"username"`
	Bio           string    `json:"bio"`
	ProfileImage  string    `json:"profile_image"`
	CreatedAt     time.Time `json:"created_at"`
}

// Recommendation represents an AI-generated NFT recommendation
type Recommendation struct {
	NFTID      string  `json:"nft_id"`
	UserID     string  `json:"user_id"`
	Score      float64 `json:"score"`
	Reason     string  `json:"reason"`
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Printf("Warning: Migration failed (tables may already exist): %v", err)
	}

	// Initialize GoFr application
	app := gofr.New()

	// Health check endpoint
	app.GET("/health", func(ctx *gofr.Context) (interface{}, error) {
		return map[string]interface{}{
			"status": "healthy",
			"service": "NFTGenie Backend",
			"timestamp": time.Now().Unix(),
		}, nil
	})

	// NFT endpoints
	app.GET("/api/nfts", getAllNFTs)
	app.GET("/api/nfts/{id}", getNFTByID)
	app.POST("/api/nfts/mint", mintNFT)
	app.GET("/api/nfts/user/{address}", getUserNFTs)

	// User endpoints
	app.POST("/api/users/connect", connectWallet)
	app.GET("/api/users/{address}", getUserProfile)
	app.PUT("/api/users/{address}", updateUserProfile)

	// AI Recommendation endpoints
	app.GET("/api/recommendations/{userId}", getRecommendations)
	app.POST("/api/recommendations/train", trainRecommendationModel)

	// Marketplace endpoints
	app.POST("/api/marketplace/list", listNFTForSale)
	app.POST("/api/marketplace/buy", buyNFT)
	app.GET("/api/marketplace/listings", getMarketplaceListings)

	// Analytics endpoints
	app.GET("/api/analytics/trending", getTrendingNFTs)
	app.GET("/api/analytics/stats", getMarketplaceStats)

	// Start server on port 8000
	app.Start()
}

// NFT Handlers
func getAllNFTs(ctx *gofr.Context) (interface{}, error) {
	nftRepo := repository.NewNFTRepository()
	
	// Get pagination params
	limit := 20
	offset := 0
	if l := ctx.Param("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}
	if o := ctx.Param("offset"); o != "" {
		fmt.Sscanf(o, "%d", &offset)
	}
	
	nfts, total, err := nftRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"nfts":  nfts,
		"total": total,
		"limit": limit,
		"offset": offset,
	}, nil
}

func getNFTByID(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid NFT ID")
	}
	
	nftRepo := repository.NewNFTRepository()
	nft, err := nftRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	return nft, nil
}

func mintNFT(ctx *gofr.Context) (interface{}, error) {
	var mintRequest struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		ImageURL    string   `json:"image_url"`
		Creator     string   `json:"creator"`
		Chain       string   `json:"chain"`
		Tags        []string `json:"tags"`
	}

	if err := ctx.Bind(&mintRequest); err != nil {
		return nil, err
	}

	// Get or create user
	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByWalletAddress(mintRequest.Creator)
	if err != nil {
		return nil, err
	}
	if user == nil {
		// Create new user
		user = models.NewUser(mintRequest.Creator)
		if err := userRepo.Create(user); err != nil {
			return nil, err
		}
	}

	// Initialize Verbwire service
	vw := services.NewVerbwireService()

	// Build request
	req := services.MintNFTRequest{
		Name:            mintRequest.Name,
		Description:     mintRequest.Description,
		ImageURL:        mintRequest.ImageURL,
		RecipientAddress: mintRequest.Creator,
		Chain:           vw.Chain,
		Quantity:        1,
	}

	// Call Verbwire API
	res, err := vw.QuickMintNFT(req)
	if err != nil {
		ctx.Logger.Errorf("mint error: %v", err)
		return map[string]interface{}{
			"success": false,
			"message": err.Error(),
		}, nil
	}

	// Save NFT to database
	nft := models.NewNFT(mintRequest.Name, mintRequest.ImageURL, user.ID, user.ID)
	nft.Description = &mintRequest.Description
	nft.ContractAddress = &res.ContractAddress
	nft.TokenID = &res.TokenID
	nft.TransactionHash = &res.TransactionHash
	nft.Chain = vw.Chain
	now := time.Now()
	nft.MintedAt = &now
	if len(mintRequest.Tags) > 0 {
		nft.Tags = mintRequest.Tags
	}

	nftRepo := repository.NewNFTRepository()
	if err := nftRepo.Create(nft); err != nil {
		ctx.Logger.Errorf("failed to save NFT: %v", err)
	}

	return map[string]interface{}{
		"success":           true,
		"message":           "NFT minted successfully",
		"nft_id":            nft.ID,
		"transaction_hash":  res.TransactionHash,
		"contract_address":  res.ContractAddress,
		"token_id":          res.TokenID,
		"opensea_url":       res.OpenseaURL,
		"chain":             vw.Chain,
	}, nil
}

func getUserNFTs(ctx *gofr.Context) (interface{}, error) {
	address := ctx.PathParam("address")
	
	// Get user by wallet address
	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByWalletAddress(address)
	if err != nil || user == nil {
		return []models.NFT{}, nil // Return empty array if user not found
	}
	
	// Get NFTs owned by user
	nftRepo := repository.NewNFTRepository()
	nfts, err := nftRepo.GetByOwner(user.ID)
	if err != nil {
		return nil, err
	}
	
	return nfts, nil
}

// User Handlers
func connectWallet(ctx *gofr.Context) (interface{}, error) {
	var walletRequest struct {
		Address   string `json:"address"`
		Signature string `json:"signature"`
		Message   string `json:"message"`
	}

	if err := ctx.Bind(&walletRequest); err != nil {
		return nil, err
	}

	// Get or create user
	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByWalletAddress(walletRequest.Address)
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}
	
	if user == nil {
		// Create new user
		user = models.NewUser(walletRequest.Address)
		if err := userRepo.CreateOrUpdate(user); err != nil {
			return nil, err
		}
	}

	// TODO: Implement actual signature verification
	// For now, we'll accept any connection
	
	return map[string]interface{}{
		"success": true,
		"user":    user,
		"token":   "jwt_token_here", // TODO: Implement JWT generation
	}, nil
}

func getUserProfile(ctx *gofr.Context) (interface{}, error) {
	address := ctx.PathParam("address")
	
	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByWalletAddress(address)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	
	// Get user stats
	stats, err := userRepo.GetUserStats(user.ID)
	if err != nil {
		ctx.Logger.Errorf("failed to get user stats: %v", err)
		stats = make(map[string]interface{})
	}
	
	return map[string]interface{}{
		"user":  user,
		"stats": stats,
	}, nil
}

func updateUserProfile(ctx *gofr.Context) (interface{}, error) {
	address := ctx.PathParam("address")
	var updateRequest struct {
		Username     string `json:"username"`
		Bio          string `json:"bio"`
		ProfileImage string `json:"profile_image"`
		Email        string `json:"email"`
	}

	if err := ctx.Bind(&updateRequest); err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByWalletAddress(address)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check username availability
	if updateRequest.Username != "" && (user.Username == nil || *user.Username != updateRequest.Username) {
		exists, err := userRepo.UsernameExists(updateRequest.Username)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("username already taken")
		}
		user.Username = &updateRequest.Username
	}

	if updateRequest.Bio != "" {
		user.Bio = &updateRequest.Bio
	}
	if updateRequest.ProfileImage != "" {
		user.ProfileImage = &updateRequest.ProfileImage
	}
	if updateRequest.Email != "" {
		user.Email = &updateRequest.Email
	}

	if err := userRepo.Update(user); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"success": true,
		"message": "Profile updated successfully",
		"user":    user,
	}, nil
}

// AI Recommendation Handlers
func getRecommendations(ctx *gofr.Context) (interface{}, error) {
	userIdStr := ctx.PathParam("userId")
	userID, err := uuid.Parse(userIdStr)
	if err != nil {
		// Try to get by wallet address
		userRepo := repository.NewUserRepository()
		user, err := userRepo.GetByWalletAddress(userIdStr)
		if err != nil || user == nil {
			return nil, fmt.Errorf("user not found")
		}
		userID = user.ID
	}
	
	// Get recommended NFTs
	nftRepo := repository.NewNFTRepository()
	nfts, err := nftRepo.GetRecommendedForUser(userID, 10)
	if err != nil {
		return nil, err
	}
	
	// Format as recommendations
	recommendations := make([]map[string]interface{}, len(nfts))
	for i, nft := range nfts {
		recommendations[i] = map[string]interface{}{
			"nft":    nft,
			"score":  0.85 + float64(10-i)*0.015, // Simulated score
			"reason": generateRecommendationReason(i),
		}
	}
	
	return recommendations, nil
}

func generateRecommendationReason(index int) string {
	reasons := []string{
		"Based on your interest in digital art",
		"Similar to NFTs you've viewed",
		"From creators you follow",
		"Trending in your favorite categories",
		"Matches your collection style",
		"Popular with similar collectors",
		"New from verified creators",
		"Rising in value recently",
		"Limited edition opportunity",
		"Recommended by the algorithm",
	}
	if index < len(reasons) {
		return reasons[index]
	}
	return "Personalized for you"
}

func trainRecommendationModel(ctx *gofr.Context) (interface{}, error) {
	// TODO: Implement model training logic
	return map[string]interface{}{
		"success": true,
		"message": "Recommendation model training initiated",
		"job_id":  "train_job_123",
	}, nil
}

// Marketplace Handlers
func listNFTForSale(ctx *gofr.Context) (interface{}, error) {
	var listingRequest struct {
		NFTID  string  `json:"nft_id"`
		Price  float64 `json:"price"`
		Seller string  `json:"seller"`
	}

	if err := ctx.Bind(&listingRequest); err != nil {
		return nil, err
	}

	// Parse NFT ID
	nftID, err := uuid.Parse(listingRequest.NFTID)
	if err != nil {
		return nil, fmt.Errorf("invalid NFT ID")
	}

	// Get seller user
	userRepo := repository.NewUserRepository()
	seller, err := userRepo.GetByWalletAddress(listingRequest.Seller)
	if err != nil || seller == nil {
		return nil, fmt.Errorf("seller not found")
	}

	// Verify NFT ownership
	nftRepo := repository.NewNFTRepository()
	nft, err := nftRepo.GetByID(nftID)
	if err != nil {
		return nil, err
	}
	if nft.OwnerID != seller.ID {
		return nil, fmt.Errorf("you don't own this NFT")
	}

	// Create listing
	listing := models.NewMarketplaceListing(nftID, seller.ID, listingRequest.Price)
	
	// TODO: Save listing to database when marketplace repository is created
	
	return map[string]interface{}{
		"success":    true,
		"listing_id": listing.ID,
		"message":    "NFT listed successfully",
	}, nil
}

func buyNFT(ctx *gofr.Context) (interface{}, error) {
	var purchaseRequest struct {
		NFTID  string `json:"nft_id"`
		Buyer  string `json:"buyer"`
		Price  string `json:"price"`
	}

	if err := ctx.Bind(&purchaseRequest); err != nil {
		return nil, err
	}

	// TODO: Process NFT purchase
	return map[string]interface{}{
		"success": true,
		"transaction_hash": "0xdef456...",
		"message": "NFT purchased successfully",
	}, nil
}

func getMarketplaceListings(ctx *gofr.Context) (interface{}, error) {
	// TODO: Get active marketplace listings
	return []map[string]interface{}{
		{
			"id":     "listing_1",
			"nft_id": "nft_1",
			"price":  "0.5",
			"seller": "0x789...",
			"status": "active",
		},
	}, nil
}

// Analytics Handlers
func getTrendingNFTs(ctx *gofr.Context) (interface{}, error) {
	nftRepo := repository.NewNFTRepository()
	trendingNFTs, err := nftRepo.GetTrending(10)
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"trending": trendingNFTs,
		"period":   "7_days",
	}, nil
}

func getMarketplaceStats(ctx *gofr.Context) (interface{}, error) {
	// Get statistics from database
	stats := make(map[string]interface{})
	
	// Get counts
	var counts struct {
		TotalNFTs   int `db:"total_nfts"`
		TotalUsers  int `db:"total_users"`
	}
	
	query := `
		SELECT 
			(SELECT COUNT(*) FROM nfts) as total_nfts,
			(SELECT COUNT(*) FROM users) as total_users`
	
	err := database.DB.Get(&counts, query)
	if err != nil {
		ctx.Logger.Errorf("failed to get stats: %v", err)
	}
	
	stats["total_nfts"] = counts.TotalNFTs
	stats["total_users"] = counts.TotalUsers
	stats["total_volume"] = "0" // TODO: Calculate from transactions
	stats["active_listings"] = 0 // TODO: Get from marketplace_listings
	stats["24h_volume"] = "0"    // TODO: Calculate from recent transactions
	stats["chain"] = "polygonAmoy"
	
	return stats, nil
}
