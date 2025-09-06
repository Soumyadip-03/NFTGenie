Write-Host "Starting NFTGenie Backend Server..." -ForegroundColor Green
Write-Host "Go installation detected: " -NoNewline
& "C:\Program Files\Go\bin\go.exe" version

cd backend
Write-Host "`nBuilding backend..." -ForegroundColor Yellow
$env:CGO_ENABLED="0"
& "C:\Program Files\Go\bin\go.exe" build -o server.exe

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful!" -ForegroundColor Green
    Write-Host "`nStarting server on http://localhost:8000" -ForegroundColor Cyan
    Write-Host "Press Ctrl+C to stop the server`n" -ForegroundColor Yellow
    .\server.exe
} else {
    Write-Host "Build failed!" -ForegroundColor Red
}
