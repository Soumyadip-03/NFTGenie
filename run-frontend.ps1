Write-Host "Starting NFTGenie Frontend..." -ForegroundColor Green

cd frontend
Write-Host "`nInstalling dependencies (if needed)..." -ForegroundColor Yellow
npm install

Write-Host "`nStarting Next.js dev server on http://localhost:3000" -ForegroundColor Cyan
Write-Host "Press Ctrl+C to stop the server`n" -ForegroundColor Yellow

$env:NEXT_TELEMETRY_DISABLED="1"
npm run dev
