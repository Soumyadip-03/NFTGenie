# NFTGenie - Dev scripts

$env:NEXT_TELEMETRY_DISABLED="1"
$env:NODE_OPTIONS="--max-old-space-size=4096"

echo "Starting frontend on http://localhost:3000"
npm run dev
