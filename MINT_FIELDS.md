# NFTGenie - NFT Minting Fields Documentation

## API Endpoint
```
POST http://localhost:8000/api/nfts/mint
Content-Type: application/json
```

## Required Fields

### 1. **name** (string, required)
- **Description**: The name of your NFT
- **Example**: `"Cosmic Dragon #001"`
- **Constraints**: Cannot be empty

### 2. **description** (string, required)
- **Description**: Detailed description of your NFT
- **Example**: `"A majestic cosmic dragon soaring through nebulas"`
- **Constraints**: Should be descriptive and engaging

### 3. **image_url** (string, required)
- **Description**: URL to the NFT image/media
- **Example**: `"https://ipfs.io/ipfs/QmXxx..."` or `"https://example.com/image.png"`
- **Constraints**: Must be a valid URL, preferably IPFS or permanent storage

### 4. **creator** (string, required)
- **Description**: Wallet address of the NFT creator
- **Example**: `"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
- **Constraints**: Must be a valid Ethereum/Polygon wallet address

## Optional Fields

### 5. **chain** (string, optional)
- **Description**: Blockchain network for minting
- **Default**: `"polygonAmoy"` (from .env configuration)
- **Options**: 
  - `"polygonAmoy"` (Polygon testnet)
  - `"polygon"` (Polygon mainnet)
  - `"ethereum"` (Ethereum mainnet)
  - `"sepolia"` (Ethereum testnet)
- **Example**: `"polygonAmoy"`

### 6. **tags** (array of strings, optional)
- **Description**: Tags for categorization and discovery
- **Example**: `["art", "digital", "fantasy", "dragon"]`
- **Use Cases**: 
  - AI recommendations
  - Search and filtering
  - Category organization

## Complete Request Examples

### Example 1: Basic NFT Mint
```json
{
  "name": "Sunset Landscape",
  "description": "A beautiful digital painting of a sunset over mountains",
  "image_url": "https://ipfs.io/ipfs/QmXxxYourImageHash",
  "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
}
```

### Example 2: NFT with Tags
```json
{
  "name": "CyberPunk City #42",
  "description": "Neon-lit futuristic cityscape with flying vehicles",
  "image_url": "https://ipfs.io/ipfs/QmCyberPunkImageHash",
  "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "tags": ["cyberpunk", "city", "futuristic", "neon", "digital-art"]
}
```

### Example 3: NFT with Custom Chain
```json
{
  "name": "Rare Collectible #001",
  "description": "First edition of the rare collectibles series",
  "image_url": "https://ipfs.io/ipfs/QmRareCollectibleHash",
  "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "chain": "polygon",
  "tags": ["rare", "collectible", "first-edition", "exclusive"]
}
```

## Successful Response
```json
{
  "success": true,
  "message": "NFT minted successfully",
  "nft_id": "123e4567-e89b-12d3-a456-426614174000",
  "transaction_hash": "0x123abc...",
  "contract_address": "0xContractAddress...",
  "token_id": "1",
  "opensea_url": "https://opensea.io/assets/matic/0x.../1",
  "chain": "polygonAmoy"
}
```

## Error Response
```json
{
  "success": false,
  "message": "Error description here"
}
```

## cURL Example
```bash
curl -X POST http://localhost:8000/api/nfts/mint \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My First NFT",
    "description": "This is my first NFT on NFTGenie",
    "image_url": "https://example.com/my-nft-image.jpg",
    "creator": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "tags": ["first", "test", "nftgenie"]
  }'
```

## JavaScript/Fetch Example
```javascript
const mintNFT = async () => {
  const response = await fetch('http://localhost:8000/api/nfts/mint', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      name: "Amazing Artwork",
      description: "A stunning piece of digital art",
      image_url: "https://ipfs.io/ipfs/QmYourHash",
      creator: walletAddress, // Get from MetaMask
      tags: ["art", "digital", "original"]
    })
  });
  
  const result = await response.json();
  console.log(result);
};
```

## Important Notes

1. **Image Storage**: 
   - Use IPFS for decentralized storage
   - Or use reliable image hosting services
   - Ensure the URL is permanently accessible

2. **Wallet Address**:
   - Must be the actual connected wallet address
   - The NFT will be minted to this address
   - Ensure MetaMask or wallet is connected

3. **Gas Fees**:
   - Minting requires gas fees on the blockchain
   - Polygon has lower fees than Ethereum
   - Ensure wallet has sufficient native tokens (MATIC/ETH)

4. **Tags Best Practices**:
   - Use relevant, descriptive tags
   - Include style tags (e.g., "3d", "pixel-art", "abstract")
   - Include category tags (e.g., "gaming", "music", "collectible")
   - Tags help with AI recommendations

5. **Validation**:
   - Name cannot be empty
   - Image URL must be valid
   - Creator must be valid wallet address
   - Description should be meaningful

## Testing with Postman

1. Open Postman
2. Create new POST request
3. URL: `http://localhost:8000/api/nfts/mint`
4. Headers: `Content-Type: application/json`
5. Body (raw JSON):
```json
{
  "name": "Test NFT",
  "description": "Testing NFT minting",
  "image_url": "https://via.placeholder.com/500",
  "creator": "0x0000000000000000000000000000000000000000",
  "tags": ["test"]
}
```

## Frontend Integration

The frontend should:
1. Connect wallet (MetaMask)
2. Get user's wallet address
3. Upload image to IPFS/storage
4. Call mint endpoint with all fields
5. Display transaction status
6. Show minted NFT details

---

**Note**: Make sure your backend is running and Verbwire API keys are configured in the .env file.
