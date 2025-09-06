# 🚀 NFTGenie Quick Start Guide

## Prerequisites Check
Before starting, ensure you have:
- ✅ PostgreSQL running (already set up!)
- ✅ Node.js installed
- ✅ Go installed
- ✅ Python with virtual environment (for AI engine)

## 🎯 Quick Start (Recommended)

### Option 1: Start Everything at Once
```powershell
.\start-all.ps1
```
This will start:
1. Backend (Go) on http://localhost:8000
2. Frontend (Next.js) on http://localhost:3000
3. AI Engine (Python) on http://localhost:5000 (optional)

### Option 2: Start Services Individually

#### Step 1: Start Backend
```powershell
cd backend
go run main.go
```
Or use the script:
```powershell
.\run-backend.ps1
```

#### Step 2: Start Frontend (in new terminal)
```powershell
cd frontend
npm run dev
```
Or use the script:
```powershell
.\run-frontend.ps1
```

#### Step 3: Start AI Engine (optional, in new terminal)
```powershell
cd ai-engine
.\venv\Scripts\Activate.ps1
python api_server_simple.py
```
Or use the script:
```powershell
.\run-ai-engine.ps1
```

## 📍 Access Points

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend** | http://localhost:3000 | Main NFTGenie web interface |
| **Backend API** | http://localhost:8000 | REST API endpoints |
| **API Health** | http://localhost:8000/health | Backend health check |
| **AI Engine** | http://localhost:5000 | AI recommendation service |
| **AI Docs** | http://localhost:5000/docs | Interactive API documentation |

## 🎨 Using NFTGenie

### 1. First Time Setup
1. Open http://localhost:3000 in your browser
2. Click "Connect Wallet" to connect MetaMask
3. Make sure you're on Polygon Amoy testnet
4. Get test MATIC from: https://faucet.polygon.technology/

### 2. Minting NFTs
1. Click on the mint form
2. Enter NFT details:
   - Name
   - Description
   - Image URL (use IPFS or direct URL)
3. Click "Mint NFT"
4. Approve transaction in MetaMask

### 3. Viewing NFTs
- Your NFTs will appear in your profile
- Browse trending NFTs on the homepage
- Get AI recommendations based on your activity

## 🛠️ Configuration

### Update API Keys (Required for Minting)
Edit `backend\.env` and add:
```env
VERBWIRE_API_KEY=your_key_here
VERBWIRE_PUBLIC_KEY=your_public_key_here
```
Get keys from: https://www.verbwire.com/

### Update WalletConnect (Required for Wallet Connection)
Edit `frontend\.env.local` and add:
```env
NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID=your_project_id
```
Get project ID from: https://cloud.walletconnect.com/

## 🔧 Troubleshooting

### Backend won't start
```powershell
# Check if PostgreSQL is running
Get-Service -Name "postgresql*"

# Check if port 8000 is in use
netstat -ano | findstr :8000

# Verify database connection
cd backend
go run test-db.go
```

### Frontend won't start
```powershell
# Check if port 3000 is in use
netstat -ano | findstr :3000

# Reinstall dependencies
cd frontend
npm install
```

### AI Engine won't start
```powershell
# Check if port 5000 is in use
netstat -ano | findstr :5000

# Reinstall Python dependencies
cd ai-engine
.\venv\Scripts\Activate.ps1
pip install -r requirements-essential.txt
```

## 📊 Checking Service Status

### Check all services
```powershell
# Backend health
Invoke-WebRequest -Uri http://localhost:8000/health

# Frontend
Invoke-WebRequest -Uri http://localhost:3000

# AI Engine
Invoke-WebRequest -Uri http://localhost:5000/
```

## 🛑 Stopping Services

### Stop all services
```powershell
.\stop-services.ps1
```

### Stop individual services
Press `Ctrl+C` in the respective terminal window

## 📱 Features Available

### Without API Keys
- ✅ Browse interface
- ✅ Connect wallet
- ✅ View mock NFTs
- ✅ Test AI recommendations
- ✅ Explore UI/UX

### With API Keys
- ✅ Mint real NFTs on Polygon Amoy
- ✅ Store NFT data in database
- ✅ Track transactions
- ✅ Full marketplace functionality

## 🎯 Development Workflow

1. **Make changes** to code
2. **Backend changes**: Restart with `go run main.go`
3. **Frontend changes**: Auto-reloads (hot reload enabled)
4. **AI Engine changes**: Restart with `python api_server_simple.py`

## 📝 Important Notes

- **Database**: PostgreSQL must be running
- **Testnet**: Use Polygon Amoy for testing
- **Gas Fees**: Need test MATIC for transactions
- **API Keys**: Required for blockchain operations

## 🚀 Ready to Start!

Run this command to start everything:
```powershell
.\start-all.ps1
```

Then open http://localhost:3000 and enjoy NFTGenie! 🎉
