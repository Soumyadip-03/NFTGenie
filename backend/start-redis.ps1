# Start Redis Server Script
Write-Host "Starting Redis Server..." -ForegroundColor Green

# Get the script directory
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$redisPath = Join-Path $scriptDir "redis\redis-server.exe"
$redisConfig = Join-Path $scriptDir "redis\redis.windows.conf"

# Check if Redis executable exists
if (Test-Path $redisPath) {
    Write-Host "Found Redis at: $redisPath" -ForegroundColor Cyan
    
    # Start Redis server
    Write-Host "Starting Redis on port 6379..." -ForegroundColor Yellow
    Write-Host "Redis will run in this window. Press Ctrl+C to stop." -ForegroundColor Cyan
    Write-Host "" 
    
    # Run Redis directly in this console (not as a separate process)
    & $redisPath $redisConfig
} else {
    Write-Host "Redis not found at: $redisPath" -ForegroundColor Red
    Write-Host "Please ensure Redis is properly installed in the 'redis' folder." -ForegroundColor Red
    exit 1
}
