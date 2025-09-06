# Start Redis Server in Background
Write-Host "Starting Redis Server in background..." -ForegroundColor Green

# Get the script directory
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$redisPath = Join-Path $scriptDir "redis\redis-server.exe"
$redisConfig = Join-Path $scriptDir "redis\redis.windows.conf"

# Check if Redis executable exists
if (Test-Path $redisPath) {
    Write-Host "Found Redis at: $redisPath" -ForegroundColor Cyan
    
    # Check if Redis is already running
    $existingRedis = Get-Process redis-server -ErrorAction SilentlyContinue
    if ($existingRedis) {
        Write-Host "Redis is already running (PID: $($existingRedis.Id))" -ForegroundColor Yellow
        exit 0
    }
    
    # Start Redis server as a hidden background process
    Write-Host "Starting Redis on port 6379..." -ForegroundColor Yellow
    $proc = Start-Process -FilePath $redisPath -ArgumentList "`"$redisConfig`"" -WindowStyle Hidden -PassThru
    
    # Wait a moment for Redis to start
    Start-Sleep -Seconds 2
    
    # Test connection
    $testConnection = & "$scriptDir\redis\redis-cli.exe" ping 2>$null
    if ($testConnection -eq "PONG") {
        Write-Host "✅ Redis is running successfully!" -ForegroundColor Green
        Write-Host "   Process ID: $($proc.Id)" -ForegroundColor Gray
        Write-Host "   Port: 6379" -ForegroundColor Gray
        Write-Host "" 
        Write-Host "Redis is running in the background." -ForegroundColor Cyan
        Write-Host "To stop Redis, use: Stop-Process -Name redis-server" -ForegroundColor Gray
    } else {
        Write-Host "Redis started but connection test failed. Please check the logs." -ForegroundColor Yellow
    }
} else {
    Write-Host "Redis not found at: $redisPath" -ForegroundColor Red
    Write-Host "Please ensure Redis is properly installed in the 'redis' folder." -ForegroundColor Red
    exit 1
}
