# Backend API

- GET    /health
- GET    /api/nfts
- GET    /api/nfts/{id}
- POST   /api/nfts/mint
- GET    /api/nfts/user/{address}

- POST   /api/users/connect
- GET    /api/users/{address}
- PUT    /api/users/{address}

- GET    /api/recommendations/{userId}
- POST   /api/recommendations/train

- POST   /api/marketplace/list
- POST   /api/marketplace/buy
- GET    /api/marketplace/listings

- GET    /api/analytics/trending
- GET    /api/analytics/stats

## Setup

1) Install Go 1.21+
2) Copy .env and set values
3) Run: go mod tidy && go run main.go (after installing Go)

## Notes
- Uses Verbwire Quick Mint on Polygon Amoy
- Requires MetaMask wallet address as recipient
- This is a testnet-only demo
