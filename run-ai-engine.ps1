Write-Host "Starting NFTGenie AI Engine..." -ForegroundColor Green
Write-Host ""

# Check if Python is installed
$python = Get-Command python -ErrorAction SilentlyContinue
if (-not $python) {
    Write-Host "Error: Python is not installed!" -ForegroundColor Red
    Write-Host "Please install Python 3.9+ from https://www.python.org/" -ForegroundColor Yellow
    exit 1
}

Write-Host "Python version: $(python --version)" -ForegroundColor Cyan
Write-Host ""

# Navigate to AI engine directory
cd ai-engine

# Check if virtual environment exists
if (-not (Test-Path "venv")) {
    Write-Host "Virtual environment not found. Creating..." -ForegroundColor Yellow
    python -m venv venv
    Write-Host "Virtual environment created." -ForegroundColor Green
}

# Activate virtual environment
Write-Host "Activating virtual environment..." -ForegroundColor Yellow
& .\venv\Scripts\Activate.ps1

# Check if dependencies are installed
$fastapi = pip show fastapi 2>$null
if (-not $fastapi) {
    Write-Host "Installing dependencies..." -ForegroundColor Yellow
    if (Test-Path "requirements-essential.txt") {
        pip install -r requirements-essential.txt
    } else {
        pip install fastapi uvicorn pydantic numpy pandas scikit-learn psycopg2-binary python-dotenv redis loguru
    }
    Write-Host "Dependencies installed." -ForegroundColor Green
}

Write-Host ""
Write-Host "Starting AI Engine server..." -ForegroundColor Cyan
Write-Host "The AI Engine will be available at: http://localhost:5000" -ForegroundColor Green
Write-Host "API Documentation: http://localhost:5000/docs" -ForegroundColor Green
Write-Host ""
Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
Write-Host ""

# Run the AI engine
python api_server.py
