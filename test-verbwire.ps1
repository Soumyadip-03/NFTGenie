# Test Verbwire API Key
$apiKey = "sk_live_6594b884-3a28-47db-a5c1-0bf497e15ddc"
$baseUrl = "https://api.verbwire.com/v1"

Write-Host "`n🔍 TESTING VERBWIRE API KEY..." -ForegroundColor Cyan
Write-Host "=" * 60 -ForegroundColor Gray

# Test 1: Simple API test
Write-Host "`n1. Testing basic API access..." -ForegroundColor Yellow
$headers = @{
    "X-API-Key" = $apiKey
    "Accept" = "application/json"
}

try {
    # Try to get chain info (simple GET request)
    $response = Invoke-RestMethod -Uri "$baseUrl/nft/chain/list" -Method GET -Headers $headers -ErrorAction Stop
    Write-Host "   ✅ API Key is valid!" -ForegroundColor Green
    Write-Host "   Available chains:" -ForegroundColor White
    $response | ConvertTo-Json -Depth 3 | Write-Host
} catch {
    $statusCode = $_.Exception.Response.StatusCode.value__
    Write-Host "   ❌ API Key test failed!" -ForegroundColor Red
    Write-Host "   Status Code: $statusCode" -ForegroundColor White
    Write-Host "   Error: $_" -ForegroundColor Gray
    
    if ($statusCode -eq 403) {
        Write-Host "`n   🔴 403 Forbidden - API Key is invalid or expired" -ForegroundColor Red
        Write-Host "   Please check your Verbwire account at: https://www.verbwire.com/dashboard" -ForegroundColor Yellow
    }
}

Write-Host ("`n" + ("=" * 60)) -ForegroundColor Gray
Write-Host "📝 NEXT STEPS:" -ForegroundColor Magenta
Write-Host "1. If API key is invalid, get a new one from:" -ForegroundColor White
Write-Host "   https://www.verbwire.com/dashboard" -ForegroundColor Cyan
Write-Host "2. Update the .env file with the new key" -ForegroundColor White
Write-Host "3. Restart the backend" -ForegroundColor White
