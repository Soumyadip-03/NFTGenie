#!/usr/bin/env pwsh

Write-Host "Setting up NFTGenie Database..." -ForegroundColor Green

# Check if PostgreSQL is running
$pgService = Get-Service -Name "postgresql*" -ErrorAction SilentlyContinue
if (-not $pgService -or $pgService.Status -ne "Running") {
    Write-Host "PostgreSQL service not found or not running. Please start PostgreSQL first." -ForegroundColor Red
    exit 1
}

# Create database
Write-Host "Creating database 'nftgenie'..." -ForegroundColor Yellow
$env:PGPASSWORD = "210806"
psql -U postgres -h localhost -c "CREATE DATABASE nftgenie;" 2>$null

if ($LASTEXITCODE -eq 0) {
    Write-Host "Database 'nftgenie' created successfully!" -ForegroundColor Green
} else {
    Write-Host "Database might already exist, continuing..." -ForegroundColor Yellow
}

# Run schema
Write-Host "Running database schema..." -ForegroundColor Yellow
psql -U postgres -h localhost -d nftgenie -f "backend/database/schema.sql"

if ($LASTEXITCODE -eq 0) {
    Write-Host "Database schema applied successfully!" -ForegroundColor Green
    Write-Host "Database setup complete! You can now start the backend." -ForegroundColor Green
} else {
    Write-Host "Failed to apply schema. Check the error above." -ForegroundColor Red
    exit 1
}