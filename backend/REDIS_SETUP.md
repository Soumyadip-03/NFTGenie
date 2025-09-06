# Redis Setup for NFTGenie Backend

## Installation Complete ✅

Redis has been successfully installed and configured for your NFTGenie backend.

## What was done:

1. **Downloaded Redis for Windows** (v3.2.100)
   - Located in: `backend/redis/`
   
2. **Created startup scripts:**
   - `start-redis.ps1` - Starts Redis server
   - `run-backend.ps1` - Runs backend with CGO disabled
   
3. **Fixed configuration issues:**
   - Added missing `APP_NAME` and `APP_VERSION` to `.env`
   - Configured Redis connection settings
   
4. **Updated main scripts:**
   - `start-all.ps1` now includes Redis startup
   - `stop-all.ps1` now properly stops Redis

## How to use:

### Start Redis only:
```powershell
cd backend
.\start-redis.ps1
```

### Start Backend only (with CGO disabled):
```powershell
cd backend
.\run-backend.ps1
```

### Start everything (recommended):
```powershell
# From project root
.\start-all.ps1
```

### Stop everything:
```powershell
# From project root
.\stop-all.ps1
```

## Redis Configuration:
- **Host:** localhost
- **Port:** 6379
- **Password:** (none - local development)
- **Config file:** `redis/redis.windows.conf`

## Troubleshooting:

### If Redis fails to start:
1. Check if port 6379 is already in use
2. Make sure no other Redis instance is running
3. Try running `.\stop-all.ps1` then start again

### If backend can't connect to Redis:
1. Ensure Redis is running first
2. Check `.env` file has correct Redis settings
3. Verify firewall isn't blocking port 6379

### CGO Error Resolution:
The backend is configured to run with `CGO_ENABLED=0` to avoid C compiler issues on Windows.

## Note:
Redis data is stored in the `redis/` directory and will persist between restarts.
