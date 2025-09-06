Write-Host "Starting NFTGenie Services..." -ForegroundColor Green
Write-Host ""

# Check prerequisites
Write-Host "Checking prerequisites..." -ForegroundColor Yellow

# Check Node.js
$node = Get-Command node -ErrorAction SilentlyContinue
if ($node) {
    Write-Host "  OK - Node.js found: $(node --version)" -ForegroundColor Green
} else {
    Write-Host "  ERROR - Node.js not found. Please install Node.js first." -ForegroundColor Red
    exit 1
}

# Check Go
$go = Get-Command go -ErrorAction SilentlyContinue
if ($go) {
    Write-Host "  OK - Go found: $(go version)" -ForegroundColor Green
} else {
    Write-Host "  ERROR - Go not found. Please install Go first." -ForegroundColor Red
    exit 1
}

# Check Redis
$redisPath = "$PWD\backend\redis\redis-server.exe"
if (Test-Path $redisPath) {
    Write-Host "  OK - Redis found" -ForegroundColor Green
} else {
    Write-Host "  ERROR - Redis not found. Please run the backend setup first." -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "Starting services..." -ForegroundColor Cyan
Write-Host ""

# Start Redis
Write-Host "1. Starting Redis Server..." -ForegroundColor Yellow
# Start Redis in background (no window needed)
& "$PWD\backend\start-redis-background.ps1"
Start-Sleep -Seconds 2

# Start Backend
Write-Host "2. Starting Backend Server..." -ForegroundColor Yellow
$backendPath = "$PWD\backend"
Start-Process PowerShell -ArgumentList "-NoExit", "-Command", "cd '$backendPath'; Write-Host 'NFTGenie Backend' -ForegroundColor Cyan; `$env:CGO_ENABLED='0'; go run main.go" -WindowStyle Normal

Start-Sleep -Seconds 3

# Start Frontend
Write-Host "3. Starting Frontend Server..." -ForegroundColor Yellow
$frontendPath = "$PWD\frontend"
Start-Process PowerShell -ArgumentList "-NoExit", "-Command", "cd '$frontendPath'; Write-Host 'NFTGenie Frontend' -ForegroundColor Cyan; `$env:NODE_NO_WARNINGS='1'; npm run dev" -WindowStyle Normal

Start-Sleep -Seconds 2

# Optional: Start AI Engine
$startAI = Read-Host "Do you want to start the AI Engine? (y/n)"
if ($startAI -eq 'y') {
    Write-Host "4. Starting AI Engine..." -ForegroundColor Yellow
    $aiPath = "$PWD\ai-engine"
    Start-Process PowerShell -ArgumentList "-NoExit", "-Command", "cd '$aiPath'; Write-Host 'NFTGenie AI Engine' -ForegroundColor Cyan; .\venv\Scripts\Activate.ps1; python api_server_simple.py" -WindowStyle Normal
}

Write-Host ""
Write-Host "Services starting up!" -ForegroundColor Green
Write-Host ""
Write-Host "Access points:" -ForegroundColor Cyan
Write-Host "  Frontend:  http://localhost:3000" -ForegroundColor White
Write-Host "  Backend:   http://localhost:8000" -ForegroundColor White
Write-Host "  AI Engine: http://localhost:5000" -ForegroundColor White
Write-Host "  Redis:     localhost:6379" -ForegroundColor White
Write-Host ""
Write-Host "To stop all services, run: .\stop-all.ps1" -ForegroundColor Yellow
