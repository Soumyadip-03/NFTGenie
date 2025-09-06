Write-Host "NFTGenie Database Setup" -ForegroundColor Green
Write-Host "======================" -ForegroundColor Green
Write-Host ""

# PostgreSQL configuration
$PSQL = "C:\Program Files\PostgreSQL\17\bin\psql.exe"

# Check if PostgreSQL is running
Write-Host "Checking PostgreSQL service..." -ForegroundColor Yellow
$pgService = Get-Service -Name "postgresql*" -ErrorAction SilentlyContinue
if ($pgService -and $pgService.Status -eq "Running") {
    Write-Host "[OK] PostgreSQL is running" -ForegroundColor Green
} else {
    Write-Host "Starting PostgreSQL service..." -ForegroundColor Yellow
    try {
        Start-Service -Name "postgresql-x64-17" -ErrorAction Stop
        Write-Host "[OK] PostgreSQL started" -ForegroundColor Green
    } catch {
        Write-Host "[WARNING] Could not start PostgreSQL service automatically" -ForegroundColor Yellow
        Write-Host "Please ensure PostgreSQL is running" -ForegroundColor Yellow
    }
}

Write-Host ""
Write-Host "Database credentials:" -ForegroundColor Cyan
Write-Host "Please enter your PostgreSQL password for user 'postgres'" -ForegroundColor Yellow
Write-Host "(This is the password you set during PostgreSQL installation)" -ForegroundColor Gray
Write-Host ""

# Get password securely
$password = Read-Host "PostgreSQL password" -AsSecureString
$BSTR = [System.Runtime.InteropServices.Marshal]::SecureStringToBSTR($password)
$pgPassword = [System.Runtime.InteropServices.Marshal]::PtrToStringAuto($BSTR)
$env:PGPASSWORD = $pgPassword

Write-Host ""
Write-Host "Step 1: Creating database 'nftgenie'..." -ForegroundColor Yellow

# Create database (ignore error if exists)
& $PSQL -U postgres -h localhost -c "CREATE DATABASE nftgenie;" 2>$null

# Check if database was created or already exists
$dbCheck = & $PSQL -U postgres -h localhost -t -c "SELECT 1 FROM pg_database WHERE datname='nftgenie';" 2>$null
if ($dbCheck -match "1") {
    Write-Host "[OK] Database 'nftgenie' is ready" -ForegroundColor Green
} else {
    Write-Host "[ERROR] Failed to create database" -ForegroundColor Red
    Write-Host "Please check your PostgreSQL installation and credentials" -ForegroundColor Yellow
    $env:PGPASSWORD = ""
    exit 1
}

Write-Host ""
Write-Host "Step 2: Applying database schema..." -ForegroundColor Yellow

# Apply schema
$schemaFile = "backend\database\schema.sql"
if (Test-Path $schemaFile) {
    $output = & $PSQL -U postgres -h localhost -d nftgenie -f $schemaFile 2>&1
    if ($LASTEXITCODE -eq 0 -or $output -match "already exists") {
        Write-Host "[OK] Schema applied successfully" -ForegroundColor Green
    } else {
        Write-Host "[WARNING] Some schema elements may already exist" -ForegroundColor Yellow
    }
} else {
    Write-Host "[ERROR] Schema file not found at: $schemaFile" -ForegroundColor Red
    $env:PGPASSWORD = ""
    exit 1
}

Write-Host ""
Write-Host "Step 3: Verifying tables..." -ForegroundColor Yellow

# List created tables
$tables = & $PSQL -U postgres -h localhost -d nftgenie -t -c "SELECT tablename FROM pg_tables WHERE schemaname='public' ORDER BY tablename;" 2>$null

if ($tables) {
    Write-Host "[OK] Tables created:" -ForegroundColor Green
    $tableList = $tables -split "`n" | Where-Object { $_.Trim() -ne "" }
    foreach ($table in $tableList) {
        Write-Host "  - $($table.Trim())" -ForegroundColor Gray
    }
} else {
    Write-Host "[WARNING] No tables found" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "Step 4: Creating .env file for backend..." -ForegroundColor Yellow

# Create .env file if it doesn't exist
$envFile = "backend\.env"
if (-not (Test-Path $envFile)) {
    @"
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=$pgPassword
DB_NAME=nftgenie
DB_SSLMODE=disable

# Verbwire API Configuration
# Get your API keys from https://www.verbwire.com/
VERBWIRE_API_KEY=your_verbwire_api_key_here
VERBWIRE_PUBLIC_KEY=your_verbwire_public_key_here
VERBWIRE_BASE_URL=https://api.verbwire.com/v1

# Blockchain Configuration
CHAIN=polygonAmoy

# Server Configuration
PORT=8000
HOST=localhost

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here_minimum_32_characters_long
JWT_EXPIRY=24h

# Redis Configuration (optional)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# AI Engine Configuration
AI_ENGINE_URL=http://localhost:5000
AI_ENGINE_API_KEY=your_ai_engine_api_key_here

# Feature Flags
ENABLE_MINTING=true
ENABLE_MARKETPLACE=true
ENABLE_AI_RECOMMENDATIONS=true

# Development Mode
ENV=development
DEBUG=true
"@ | Out-File -FilePath $envFile -Encoding UTF8
    Write-Host "[OK] Created .env file at: $envFile" -ForegroundColor Green
    Write-Host "  [!] Please update the API keys in the .env file" -ForegroundColor Yellow
} else {
    Write-Host "[OK] .env file already exists" -ForegroundColor Green
    # Update only the database password
    $envContent = Get-Content $envFile -Raw
    $envContent = $envContent -replace "DB_PASSWORD=.*", "DB_PASSWORD=$pgPassword"
    $envContent | Out-File -FilePath $envFile -Encoding UTF8 -NoNewline
    Write-Host "  Updated database password in .env" -ForegroundColor Gray
}

# Clear password from environment
$env:PGPASSWORD = ""

Write-Host ""
Write-Host "========================================" -ForegroundColor Green
Write-Host "Database setup complete!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green
Write-Host ""
Write-Host "Database Details:" -ForegroundColor Cyan
Write-Host "  Server:   localhost:5432" -ForegroundColor White
Write-Host "  Database: nftgenie" -ForegroundColor White
Write-Host "  Username: postgres" -ForegroundColor White
Write-Host "  Password: [your password]" -ForegroundColor White
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "1. Update API keys in backend\.env file" -ForegroundColor White
Write-Host "2. Run the backend: .\run-backend.ps1" -ForegroundColor White
Write-Host "3. Run the frontend: .\run-frontend.ps1" -ForegroundColor White
Write-Host ""
Write-Host "To connect with pgAdmin or other tools:" -ForegroundColor Cyan
$connString = "postgresql://postgres:[password]@localhost:5432/nftgenie"
Write-Host "  Connection string: $connString" -ForegroundColor Gray
