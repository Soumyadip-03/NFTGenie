# NFTGenie Complete Setup Guide

## Prerequisites

### Required Software
- **Node.js** v18+ and npm
- **Go** 1.21+
- **PostgreSQL** 14+
- **Redis** 6+ (for caching)
- **Python** 3.9+
- **Git**

### Optional Software
- Docker & Docker Compose (for containerized deployment)
- ngrok (for testing webhooks locally)

## Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/nftgenie.git
cd nftgenie
```

### 2. Database Setup

#### Install PostgreSQL
```bash
# Windows (using Chocolatey)
choco install postgresql

# macOS
brew install postgresql

# Ubuntu/Debian
sudo apt update
sudo apt install postgresql postgresql-contrib
```

#### Create Database
```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE nftgenie;

# Exit
\q

# Run migrations
cd backend
psql -U postgres -d nftgenie -f database/schema.sql
```

### 3. Redis Setup
```bash
# Windows
choco install redis

# macOS
brew install redis

# Ubuntu/Debian
sudo apt install redis-server

# Start Redis
redis-server
```

### 4. Backend Setup

#### Configure Environment
```bash
cd backend
cp .env.example .env
# Edit .env with your configuration
```

#### Required .env values:
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=nftgenie

# Verbwire API (Get from https://www.verbwire.com/)
VERBWIRE_API_KEY=your_api_key
VERBWIRE_PUBLIC_KEY=your_public_key
VERBWIRE_BASE_URL=https://api.verbwire.com/v1
CHAIN=polygonAmoy

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
```

#### Install Dependencies & Run
```bash
# Install Go dependencies
go mod download

# Run the backend
go run main.go

# Or build and run
go build -o server.exe
./server.exe
```

The backend will be available at `http://localhost:8000`

### 5. Frontend Setup

#### Configure Environment
```bash
cd frontend
cp .env.local.example .env.local
# Edit .env.local with your configuration
```

#### Required .env.local values:
```env
NEXT_PUBLIC_APP_NAME=NFTGenie
NEXT_PUBLIC_API_URL=http://localhost:8000
NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID=your_project_id_from_walletconnect
```

#### Install Dependencies & Run
```bash
# Install dependencies
npm install

# Run development server
npm run dev
```

The frontend will be available at `http://localhost:3000`

### 6. AI Engine Setup

#### Create Python Virtual Environment
```bash
cd ai-engine
python -m venv venv

# Windows
venv\Scripts\activate

# macOS/Linux
source venv/bin/activate
```

#### Install Dependencies
```bash
pip install -r requirements.txt
```

#### Configure Environment
```bash
cp ../backend/.env .env
# The AI engine uses the same database configuration
```

#### Run AI Engine Server
```bash
python api_server.py
```

The AI engine will be available at `http://localhost:5000`

## Production Deployment

### Using Docker Compose

Create `docker-compose.yml`:
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:14
    environment:
      POSTGRES_DB: nftgenie
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./backend/database/schema.sql:/docker-entrypoint-initdb.d/schema.sql
    ports:
      - "5432:5432"

  redis:
    image: redis:7
    ports:
      - "6379:6379"

  backend:
    build: ./backend
    depends_on:
      - postgres
      - redis
    environment:
      - DB_HOST=postgres
      - REDIS_HOST=redis
    env_file:
      - ./backend/.env
    ports:
      - "8000:8000"

  frontend:
    build: ./frontend
    depends_on:
      - backend
    env_file:
      - ./frontend/.env.local
    ports:
      - "3000:3000"

  ai-engine:
    build: ./ai-engine
    depends_on:
      - postgres
      - redis
    environment:
      - DB_HOST=postgres
      - REDIS_HOST=redis
    env_file:
      - ./backend/.env
    ports:
      - "5000:5000"

volumes:
  postgres_data:
```

Run with:
```bash
docker-compose up -d
```

### Manual Production Setup

#### 1. Setup PostgreSQL
```bash
# Create production database
sudo -u postgres createdb nftgenie_prod
sudo -u postgres psql -d nftgenie_prod -f backend/database/schema.sql
```

#### 2. Build Backend
```bash
cd backend
CGO_ENABLED=0 GOOS=linux go build -o server
```

#### 3. Build Frontend
```bash
cd frontend
npm run build
npm run start
```

#### 4. Setup AI Engine with Gunicorn
```bash
cd ai-engine
pip install gunicorn
gunicorn -w 4 -k uvicorn.workers.UvicornWorker api_server:app --bind 0.0.0.0:5000
```

#### 5. Setup Nginx (Optional)
```nginx
server {
    listen 80;
    server_name nftgenie.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    location /api {
        proxy_pass http://localhost:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    location /ai {
        proxy_pass http://localhost:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

### AI Engine Tests
```bash
cd ai-engine
pytest
```

## Troubleshooting

### Database Connection Issues
- Ensure PostgreSQL is running: `pg_isready`
- Check credentials in `.env` file
- Verify database exists: `psql -U postgres -l`

### Verbwire API Issues
- Verify API keys are correct
- Check you're using the correct chain (polygonAmoy for testnet)
- Ensure you have testnet MATIC for gas fees

### Redis Connection Issues
- Check Redis is running: `redis-cli ping`
- Verify Redis configuration in `.env`

### Port Conflicts
- Backend default: 8000
- Frontend default: 3000
- AI Engine default: 5000
- PostgreSQL default: 5432
- Redis default: 6379

Change ports in `.env` files if needed.

## Monitoring

### Check System Health
```bash
# Backend health
curl http://localhost:8000/health

# AI engine health
curl http://localhost:5000/

# Database connections
psql -U postgres -c "SELECT count(*) FROM pg_stat_activity WHERE datname = 'nftgenie';"
```

### View Logs
```bash
# Backend logs (if using systemd)
journalctl -u nftgenie-backend -f

# Docker logs
docker-compose logs -f backend
```

## Security Checklist

- [ ] Change default database passwords
- [ ] Set strong JWT_SECRET
- [ ] Configure CORS properly for production
- [ ] Use HTTPS in production
- [ ] Enable rate limiting
- [ ] Secure API keys
- [ ] Set up firewall rules
- [ ] Enable database SSL
- [ ] Regular backups

## Support

For issues or questions:
1. Check the [FAQ](docs/FAQ.md)
2. Search [existing issues](https://github.com/yourusername/nftgenie/issues)
3. Create a new issue with details

## License

MIT License - see LICENSE file for details
