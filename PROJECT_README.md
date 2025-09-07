# NFTGenie - AI-Powered NFT Platform

## 🎯 What is NFTGenie?

NFTGenie is a modern, full-stack NFT minting and recommendation platform that combines blockchain technology with artificial intelligence. Users can create, mint, and discover unique digital assets on the Polygon network with AI-powered personalized recommendations.

## 🚀 Key Features

### ✨ **Core Functionality**
- **NFT Minting**: Create and mint NFTs on Polygon Amoy testnet
- **Wallet Integration**: MetaMask wallet connection via RainbowKit
- **AI Recommendations**: Personalized NFT suggestions based on user behavior
- **Real-time Notifications**: Live updates for minting progress and transactions
- **Modern UI**: Clean, responsive design with glassmorphism effects

### 🎨 **User Experience**
- **Notification System**: Dropdown notifications with history and clear functionality
- **Theme Consistency**: Modern blue/green color scheme throughout
- **Responsive Design**: Works on desktop, tablet, and mobile
- **Settings Panel**: User preferences and notification controls
- **Preview System**: Live NFT preview while creating

## 🔧 Quick Start Guide

### **Prerequisites**
- Node.js 18+
- Go 1.21+
- Python 3.8+
- PostgreSQL 13+
- MetaMask wallet

### **1. Clone & Setup**
```bash
git clone <repository-url>
cd NFTGenie
```

### **2. Database Setup**
```sql
-- Create database
CREATE DATABASE nftgenie;

-- Run schema
psql -U postgres -d nftgenie -f backend/database/schema.sql
```

### **3. Backend Setup**
```bash
cd backend
cp .env.example .env
# Edit .env with your database credentials and Verbwire API keys
go mod tidy
go run main.go
```

### **4. Frontend Setup**
```bash
cd frontend
npm install
npm run dev
```

### **5. AI Engine Setup**
```bash
cd ai-engine
pip install -r requirements.txt
python api_server_simple.py  # For development
```

### **6. Access the App**
- Frontend: http://localhost:3000
- Backend API: http://localhost:8000
- AI Engine: http://localhost:5000

## 🧪 Testing Your Setup

### **Test NFT Data**
```json
{
  "name": "Cosmic Dragon #001",
  "image_url": "https://via.placeholder.com/500/FF6B6B/FFFFFF?text=Cosmic+Dragon",
  "description": "A majestic cosmic dragon soaring through nebulas and stars.",
  "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7"
}
```

### **Testing Flow**
1. Open http://localhost:3000
2. Connect MetaMask wallet
3. Fill mint form with test data above
4. Click "Mint NFT"
5. Check notifications for progress
6. Verify transaction on Polygonscan

## 🏗️ Architecture Overview

```
NFTGenie/
├── frontend/          # Next.js React application
├── backend/           # Go API server
├── ai-engine/         # Python AI recommendation service
└── database/          # PostgreSQL schema and data
```

## 🛠️ Technology Stack

### **Frontend** (Next.js 15)
- **Framework**: Next.js with TypeScript
- **Styling**: Tailwind CSS with custom themes
- **Wallet**: RainbowKit + Wagmi for Web3 integration
- **State Management**: React Context API
- **Animations**: Custom CSS animations and transitions

### **Backend** (Go)
- **Framework**: GoFr framework
- **Database**: PostgreSQL with SQLX
- **Blockchain**: Verbwire API for Polygon minting
- **Environment**: dotenv for configuration
- **UUID**: Google UUID for unique identifiers

### **AI Engine** (Python)
- **Framework**: FastAPI for API server
- **ML Libraries**: scikit-learn, pandas, numpy
- **Caching**: Redis for performance
- **Database**: PostgreSQL integration
- **Models**: Collaborative filtering and content-based recommendations

### **Database** (PostgreSQL)
- **NFTs Table**: Store minted NFT metadata
- **Users Table**: User profiles and wallet addresses
- **Interactions Table**: User behavior tracking for AI

## 📁 Project Structure

```
NFTGenie/
├── frontend/
│   ├── app/
│   │   ├── page.tsx              # Main homepage
│   │   ├── settings/page.tsx     # Settings page
│   │   ├── layout.tsx            # App layout
│   │   └── globals.css           # Global styles
│   ├── components/
│   │   ├── MintForm.tsx          # NFT creation form
│   │   ├── ConnectWallet.tsx     # Wallet connection
│   │   ├── NotificationButton.tsx # Notification dropdown
│   │   ├── Settings.tsx          # Settings button
│   │   └── Logo.tsx              # App logo
│   ├── contexts/
│   │   └── NotificationContext.tsx # Global notifications
│   └── package.json
│
├── backend/
│   ├── main.go                   # Main server file
│   ├── database/
│   │   ├── db.go                 # Database connection
│   │   └── schema.sql            # Database schema
│   ├── models/
│   │   └── models.go             # Data models
│   ├── repository/
│   │   ├── nft_repository.go     # NFT data access
│   │   └── user_repository.go    # User data access
│   ├── services/
│   │   └── verbwire.go           # Blockchain service
│   ├── .env                      # Environment variables
│   └── go.mod
│
├── ai-engine/
│   ├── api_server.py             # Full AI server (production)
│   ├── api_server_simple.py      # Simple AI server (development)
│   ├── recommender.py            # ML recommendation logic
│   ├── advanced_recommender.py   # Advanced ML models
│   └── requirements.txt
│
└── docs/
    ├── TEST_MINT_DATA.md         # Testing guide
    ├── AI_ENGINE_COMPARISON.md   # AI engine comparison
    └── SETUP.md                  # Setup instructions
```

## 🌐 API Endpoints

### **Backend API** (Port 8000)
```
GET    /health                    # Health check
GET    /api/nfts                  # List all NFTs
GET    /api/nfts/{id}             # Get specific NFT
POST   /api/nfts/mint             # Mint new NFT
GET    /api/nfts/user/{address}   # Get user's NFTs
POST   /api/users/connect         # Connect wallet
GET    /api/users/{address}       # Get user profile
PUT    /api/users/{address}       # Update user profile
```

### **AI Engine API** (Port 5000)
```
POST   /recommend                 # Get personalized recommendations
POST   /train                     # Train ML model
GET    /health                    # Health check
```

## 🎨 Design System

### **Color Palette**
- **Primary**: Blue (#3b82f6)
- **Secondary**: Green (#10b981)
- **Background**: Slate (#0f172a, #1e293b)
- **Text**: Slate variations (#f1f5f9, #94a3b8)
- **Accent**: Amber (#f59e0b) for highlights

### **Components**
- **Glass Effect**: Translucent cards with backdrop blur
- **Hover Animations**: Subtle lift and glow effects
- **Gradient Text**: Blue-to-green gradient for headings
- **Notification System**: Slide-down animated dropdown
- **Responsive Grid**: Adaptive layout for all screen sizes

## 🔐 Environment Variables

### **Backend (.env)**
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=nftgenie

# Verbwire API (get from verbwire.com)
VERBWIRE_API_KEY=your_api_key
VERBWIRE_PUBLIC_KEY=your_public_key

# Server
PORT=8000
HOST=localhost

# JWT
JWT_SECRET=your_jwt_secret_32_chars_minimum
```

## 🚀 Production Deployment

### **Production Checklist**
- [ ] Set up production PostgreSQL database
- [ ] Configure Redis for AI caching
- [ ] Set environment variables
- [ ] Use production Verbwire API keys
- [ ] Enable HTTPS
- [ ] Set up monitoring and logging

### **Recommended Hosting**
- **Frontend**: Vercel, Netlify
- **Backend**: Railway, Render, DigitalOcean
- **Database**: Supabase, PlanetScale, AWS RDS
- **AI Engine**: Railway, Render, Google Cloud Run

## 📊 Features Status

### **Implemented ✅**
- NFT minting on Polygon Amoy
- Wallet connection with RainbowKit
- Real-time notifications system
- AI recommendation engine (2 versions)
- Modern responsive UI
- Settings and preferences
- Database integration
- Error handling and validation

### **Planned 🔄**
- NFT marketplace functionality
- User profiles and collections
- Social features and sharing
- Advanced AI models
- Mobile app
- Multi-chain support

## 🆘 Troubleshooting

### **Common Issues**

**Backend won't start:**
- Check PostgreSQL is running
- Verify database credentials in `.env`
- Ensure port 8000 is available

**Frontend errors:**
- Run `npm install` to install dependencies
- Check Node.js version (18+)
- Verify backend is running on port 8000

**Wallet connection fails:**
- Install MetaMask extension
- Switch to Polygon Amoy testnet
- Check wallet has test MATIC

**NFT minting fails:**
- Verify Verbwire API keys in `.env`
- Check wallet is connected
- Ensure valid image URL

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

- **Documentation**: Check `/docs` folder for detailed guides
- **Issues**: Create GitHub issues for bugs
- **Discord**: Join our community for support
- **Email**: support@nftgenie.com

## 🎉 Acknowledgments

- **Verbwire**: Blockchain infrastructure
- **RainbowKit**: Wallet connection
- **Tailwind CSS**: Styling framework
- **Next.js**: React framework
- **GoFr**: Go web framework
- **FastAPI**: Python API framework

---

**Built with ❤️ by the DEVXPERTS Team**