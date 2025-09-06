# Start Frontend without deprecation warnings
Write-Host "Starting Frontend (suppressing deprecation warnings)..." -ForegroundColor Green

# Set environment variable to suppress warnings
$env:NODE_NO_WARNINGS = "1"

# Start the development server
Write-Host "Frontend will be available at http://localhost:3000" -ForegroundColor Cyan
npm run dev
