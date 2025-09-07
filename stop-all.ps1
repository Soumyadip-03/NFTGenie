Write-Host "Stopping NFTGenie Services..." -ForegroundColor Red
Write-Host ""

# Stop Redis
Write-Host "1. Stopping Redis Server..." -ForegroundColor Yellow
$redisProcess = Get-Process redis-server -ErrorAction SilentlyContinue
if ($redisProcess) {
    Stop-Process -Name redis-server -Force
    Write-Host "   Redis stopped" -ForegroundColor Green
} else {
    Write-Host "   Redis was not running" -ForegroundColor Gray
}

# Stop Go backend processes
Write-Host "2. Stopping Backend Server..." -ForegroundColor Yellow
$goProcesses = Get-Process go -ErrorAction SilentlyContinue
if ($goProcesses) {
    Stop-Process -Name go -Force
    Write-Host "   Backend stopped" -ForegroundColor Green
} else {
    Write-Host "   Backend was not running" -ForegroundColor Gray
}

# Stop Node.js frontend processes
Write-Host "3. Stopping Frontend Server..." -ForegroundColor Yellow
$nodeProcesses = Get-Process node -ErrorAction SilentlyContinue
if ($nodeProcesses) {
    Stop-Process -Name node -Force
    Write-Host "   Frontend stopped" -ForegroundColor Green
} else {
    Write-Host "   Frontend was not running" -ForegroundColor Gray
}

# Stop Python AI engine processes
Write-Host "4. Stopping AI Engine..." -ForegroundColor Yellow
$pythonProcesses = Get-Process python -ErrorAction SilentlyContinue
if ($pythonProcesses) {
    Stop-Process -Name python -Force
    Write-Host "   AI Engine stopped" -ForegroundColor Green
} else {
    Write-Host "   AI Engine was not running" -ForegroundColor Gray
}

Write-Host ""
Write-Host "All NFTGenie services stopped!" -ForegroundColor Green