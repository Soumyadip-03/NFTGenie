Write-Host "Stopping NFTGenie Services..." -ForegroundColor Yellow
Write-Host ""

# Stop Frontend (Next.js on port 3000)
Write-Host "Checking for Frontend on port 3000..." -ForegroundColor Cyan
$frontend = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue | Select-Object -First 1
if ($frontend) {
    $pid = $frontend.OwningProcess
    Write-Host "  Found process PID: $pid" -ForegroundColor Gray
    Stop-Process -Id $pid -Force
    Write-Host "  ✓ Frontend stopped" -ForegroundColor Green
} else {
    Write-Host "  - Frontend not running" -ForegroundColor DarkGray
}

# Stop Backend (Go server on port 8000)
Write-Host "Checking for Backend on port 8000..." -ForegroundColor Cyan
$backend = Get-NetTCPConnection -LocalPort 8000 -ErrorAction SilentlyContinue | Select-Object -First 1
if ($backend) {
    $pid = $backend.OwningProcess
    Write-Host "  Found process PID: $pid" -ForegroundColor Gray
    Stop-Process -Id $pid -Force
    Write-Host "  ✓ Backend stopped" -ForegroundColor Green
} else {
    Write-Host "  - Backend not running" -ForegroundColor DarkGray
}

# Stop AI Engine (Python on port 5000)
Write-Host "Checking for AI Engine on port 5000..." -ForegroundColor Cyan
$aiengine = Get-NetTCPConnection -LocalPort 5000 -ErrorAction SilentlyContinue | Select-Object -First 1
if ($aiengine) {
    $pid = $aiengine.OwningProcess
    Write-Host "  Found process PID: $pid" -ForegroundColor Gray
    Stop-Process -Id $pid -Force
    Write-Host "  ✓ AI Engine stopped" -ForegroundColor Green
} else {
    Write-Host "  - AI Engine not running" -ForegroundColor DarkGray
}

# Stop Redis (on port 6379)
Write-Host "Checking for Redis on port 6379..." -ForegroundColor Cyan
$redis = Get-NetTCPConnection -LocalPort 6379 -ErrorAction SilentlyContinue | Select-Object -First 1
if ($redis) {
    $pid = $redis.OwningProcess
    Write-Host "  Found process PID: $pid" -ForegroundColor Gray
    Stop-Process -Id $pid -Force
    Write-Host "  ✓ Redis stopped" -ForegroundColor Green
} else {
    Write-Host "  - Redis not running" -ForegroundColor DarkGray
}

# Kill any remaining node processes
$nodeProcesses = Get-Process node -ErrorAction SilentlyContinue
if ($nodeProcesses) {
    Write-Host ""
    Write-Host "Cleaning up Node.js processes..." -ForegroundColor Yellow
    $nodeProcesses | ForEach-Object {
        Stop-Process -Id $_.Id -Force
        Write-Host "  Stopped Node process PID: $($_.Id)" -ForegroundColor Gray
    }
}

# Kill any remaining Go processes
$goProcesses = Get-Process go, server, nftgenie -ErrorAction SilentlyContinue
if ($goProcesses) {
    Write-Host ""
    Write-Host "Cleaning up Go processes..." -ForegroundColor Yellow
    $goProcesses | ForEach-Object {
        Stop-Process -Id $_.Id -Force
        Write-Host "  Stopped Go process PID: $($_.Id)" -ForegroundColor Gray
    }
}

# Kill any remaining Redis processes
$redisProcesses = Get-Process redis-server -ErrorAction SilentlyContinue
if ($redisProcesses) {
    Write-Host ""
    Write-Host "Cleaning up Redis processes..." -ForegroundColor Yellow
    $redisProcesses | ForEach-Object {
        Stop-Process -Id $_.Id -Force
        Write-Host "  Stopped Redis process PID: $($_.Id)" -ForegroundColor Gray
    }
}

Write-Host ""
Write-Host "✅ All NFTGenie services stopped!" -ForegroundColor Green
Write-Host ""
Write-Host "To start services again, use:" -ForegroundColor Cyan
Write-Host "  ./run-backend.ps1   - Start backend"
Write-Host "  ./run-frontend.ps1  - Start frontend"
Write-Host "  ./start-all.ps1     - Start all services"
