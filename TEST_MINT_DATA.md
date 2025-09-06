# NFT Minting Test Data & Verification Guide

## 🎨 Test Data for Mint Form

### Test Case 1: Basic NFT
```
NFT Name: Cosmic Dragon #001
Image URL: https://via.placeholder.com/500/FF6B6B/FFFFFF?text=Cosmic+Dragon
Description: A majestic cosmic dragon soaring through nebulas and stars. This mythical creature embodies the power of the cosmos.
```

### Test Case 2: Art NFT with IPFS
```
NFT Name: Abstract Dreams
Image URL: https://ipfs.io/ipfs/QmXoypizjW3WknFiJnKLwHCnL72vedxjQkDDP1mXWo6uco/I/m/Image_placeholder.png
Description: An abstract representation of dreams and consciousness, blending vibrant colors with ethereal shapes.
```

### Test Case 3: Gaming NFT
```
NFT Name: Legendary Sword of Ethereum
Image URL: https://picsum.photos/500/500
Description: A legendary sword forged in the depths of the blockchain. +100 attack power, +50 defense. Rare drop from the Genesis dungeon.
```

### Test Case 4: Simple Test
```
NFT Name: Test NFT 001
Image URL: https://dummyimage.com/500x500/6366f1/ffffff&text=NFT
Description: This is a test NFT to verify the minting process is working correctly.
```

---

## 🔍 How to Verify Minting is Working

### 1. **Check Browser Console**
Press `F12` → Go to Console tab

**Success Response:**
```json
{
  "success": true,
  "message": "NFT minted successfully",
  "nft_id": "123e4567-e89b-12d3-a456-426614174000",
  "transaction_hash": "0x123abc...",
  "contract_address": "0xContract...",
  "token_id": "1",
  "opensea_url": "https://opensea.io/assets/matic/0x.../1",
  "chain": "polygonAmoy"
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "Error description here"
}
```

### 2. **Check Network Tab**
Press `F12` → Network tab → Look for POST request to `/api/nfts/mint`

- **Status Code**: Should be 200 (green)
- **Response**: Check the response data
- **Request Payload**: Verify your data was sent

### 3. **Check Backend Logs**
Look at your backend PowerShell window:

**Success Log:**
```
INFO [timestamp] POST /api/nfts/mint - 200 OK
```

**Error Log:**
```
ERRO [timestamp] mint error: [error details]
```

### 4. **Visual Feedback on Frontend**
After clicking "Mint NFT", you should see:

✅ **Success:**
- Loading spinner during minting
- Success toast/notification
- Transaction hash displayed
- Link to view on Polygonscan

❌ **Error:**
- Error message displayed
- Red notification
- Specific error details

---

## 🧪 Testing Without Real Wallet

### Mock Wallet Address for Testing:
```
0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7
```

### Using Postman/cURL:

**Postman:**
1. Method: POST
2. URL: `http://localhost:8000/api/nfts/mint`
3. Headers: `Content-Type: application/json`
4. Body (raw JSON):
```json
{
  "name": "Test NFT via Postman",
  "description": "Testing the mint endpoint directly",
  "image_url": "https://via.placeholder.com/500",
  "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7",
  "tags": ["test", "postman"]
}
```

**cURL Command:**
```bash
curl -X POST http://localhost:8000/api/nfts/mint \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test NFT via cURL",
    "description": "Testing mint endpoint",
    "image_url": "https://via.placeholder.com/500",
    "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7"
  }'
```

---

## 🚦 Quick Verification Checklist

1. **Frontend Loading?**
   - [ ] Form shows loading state when clicking "Mint NFT"
   - [ ] Button becomes disabled during minting

2. **API Call Made?**
   - [ ] Network tab shows POST to `/api/nfts/mint`
   - [ ] Request has all required fields

3. **Backend Processing?**
   - [ ] Backend console shows request received
   - [ ] No error messages in backend

4. **Response Received?**
   - [ ] Frontend receives response
   - [ ] Success/error message displayed

5. **Database Check:**
   ```sql
   -- Run in PostgreSQL to verify NFT was saved
   SELECT * FROM nfts ORDER BY created_at DESC LIMIT 1;
   ```

---

## ⚠️ Common Issues & Solutions

### Issue 1: "Network Error"
- **Check**: Is backend running on port 8000?
- **Fix**: Restart backend

### Issue 2: "Invalid wallet address"
- **Check**: Is wallet connected?
- **Fix**: Connect MetaMask first

### Issue 3: "Verbwire API error"
- **Check**: Are API keys in .env correct?
- **Fix**: Verify VERBWIRE_API_KEY in backend/.env

### Issue 4: No response
- **Check**: Browser console for CORS errors
- **Fix**: Backend CORS settings

---

## 📱 Test Flow

1. Open http://localhost:3000
2. Click "Connect Wallet" (or use mock address)
3. Fill form with test data above
4. Click "Mint NFT"
5. Check console (F12)
6. Verify in backend logs
7. Check response

---

## 🎯 Expected Success Flow

1. Click "Mint NFT" → Button shows "Minting..."
2. Backend receives request
3. Verbwire API called (may fail if no real API key)
4. NFT saved to database
5. Success response returned
6. Frontend shows success message
7. Transaction details displayed
