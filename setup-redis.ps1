# Redis Setup Script for Windows
Write-Host "Setting up Redis for NFTGenie..." -ForegroundColor Green

$redisDir = "$PWD\backend\redis"
$redisUrl = "https://github.com/microsoftarchive/redis/releases/download/win-3.2.100/Redis-x64-3.2.100.zip"
$zipFile = "$redisDir\redis.zip"

# Create redis directory if it doesn't exist
if (!(Test-Path $redisDir)) {
    New-Item -ItemType Directory -Path $redisDir -Force
}

# Check if Redis is already installed
if (Test-Path "$redisDir\redis-server.exe") {
    Write-Host "Redis is already installed!" -ForegroundColor Green
    exit 0
}

Write-Host "Downloading Redis..." -ForegroundColor Yellow
try {
    Invoke-WebRequest -Uri $redisUrl -OutFile $zipFile -UseBasicParsing
    Write-Host "Download completed!" -ForegroundColor Green
} catch {
    Write-Host "Failed to download Redis: $_" -ForegroundColor Red
    exit 1
}

Write-Host "Extracting Redis..." -ForegroundColor Yellow
try {
    Expand-Archive -Path $zipFile -DestinationPath $redisDir -Force
    Remove-Item $zipFile -Force
    Write-Host "Redis extracted successfully!" -ForegroundColor Green
} catch {
    Write-Host "Failed to extract Redis: $_" -ForegroundColor Red
    exit 1
}

Write-Host "Redis setup completed!" -ForegroundColor Green
Write-Host "Redis server location: $redisDir\redis-server.exe" -ForegroundColor Cyan