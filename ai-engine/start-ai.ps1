# AI Engine Startup Script
Write-Host "NFTGenie AI Engine Startup" -ForegroundColor Cyan
Write-Host ""

# Check if database is accessible
$dbAccessible = $false
try {
    # Test database connection using psql if available
    $env:PGPASSWORD = "Soumyadip@18"
    $testCmd = "psql -h localhost -U postgres -d nftgenie -c 'SELECT 1' 2>&1"
    $result = Invoke-Expression $testCmd
    if ($result -match "1") {
        $dbAccessible = $true
    }
} catch {
    $dbAccessible = $false
}

# Activate virtual environment
Write-Host "Activating virtual environment..." -ForegroundColor Gray
& ".\venv\Scripts\Activate.ps1"

# Choose which server to run
if ($dbAccessible) {
    Write-Host "✅ Database is accessible - Starting full AI Engine" -ForegroundColor Green
    Write-Host "This will use real data from your database" -ForegroundColor Gray
    python api_server.py
} else {
    Write-Host "⚠️  Database not accessible - Starting simplified AI Engine" -ForegroundColor Yellow
    Write-Host "This will use sample data for testing" -ForegroundColor Gray
    python api_server_simple.py
}
