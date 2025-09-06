# Run Backend Script with CGO disabled
Write-Host "Starting NFTGenie Backend..." -ForegroundColor Green

# Set environment variable to disable CGO
$env:CGO_ENABLED = "0"

# Get the script directory
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path

# Check if executable exists, if not build it
$exePath = Join-Path $scriptDir "nftgenie.exe"
if (-not (Test-Path $exePath)) {
    Write-Host "Building backend executable..." -ForegroundColor Yellow
    go build -o nftgenie.exe main.go
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Failed to build backend" -ForegroundColor Red
        exit 1
    }
}

Write-Host "Starting backend server on port 8000..." -ForegroundColor Cyan
Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow

# Run the backend
& $exePath
