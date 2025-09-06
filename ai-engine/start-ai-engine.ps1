# Start AI Engine Script
Write-Host "Starting NFTGenie AI Engine..." -ForegroundColor Green

# Get the script directory
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path

# Activate virtual environment
$venvPath = Join-Path $scriptDir "venv\Scripts\Activate.ps1"
if (Test-Path $venvPath) {
    & $venvPath
    Write-Host "Virtual environment activated" -ForegroundColor Cyan
} else {
    Write-Host "Virtual environment not found!" -ForegroundColor Red
    exit 1
}

# Check if FastAPI is installed
python -c "import fastapi" 2>$null
if ($LASTEXITCODE -ne 0) {
    Write-Host "FastAPI not installed. Installing now..." -ForegroundColor Yellow
    pip install fastapi uvicorn loguru numpy scikit-learn
}

Write-Host "Starting AI Engine on port 5000..." -ForegroundColor Yellow
Write-Host "API documentation will be available at http://localhost:5000/docs" -ForegroundColor Cyan

# Run the simplified server (works without database/Redis)
python api_server_simple.py
