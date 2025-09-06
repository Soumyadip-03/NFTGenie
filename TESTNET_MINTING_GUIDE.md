# 🚀 Complete Guide: Minting NFTs on Testnet

## 📋 Prerequisites Checklist

### 1. **MetaMask Wallet Setup**
- [ ] Install MetaMask extension
- [ ] Create/Import wallet
- [ ] Switch to Polygon Amoy Testnet

### 2. **Get Test MATIC (Free)**
- [ ] Get testnet MATIC tokens
- [ ] Verify balance in wallet

### 3. **Verbwire Account Setup**
- [ ] Create account at verbwire.com
- [ ] Verify email
- [ ] Get API keys
- [ ] Add testnet credits

---

## 🦊 Step 1: Setup MetaMask for Polygon Amoy Testnet

### Add Polygon Amoy Network to MetaMask:

1. Open MetaMask
2. Click network dropdown → "Add Network" → "Add Network Manually"
3. Enter these details:

```
Network Name: Polygon Amoy Testnet
RPC URL: https://rpc-amoy.polygon.technology/
Chain ID: 80002
Currency Symbol: MATIC
Block Explorer: https://amoy.polygonscan.com/
```

### Alternative RPC URLs (if first one doesn't work):
```
https://polygon-amoy.blockpi.network/v1/rpc/public
https://polygon-amoy.gateway.tenderly.co
https://polygon-amoy.drpc.org
```

---

## 💰 Step 2: Get Free Test MATIC

### Option A: Polygon Faucet (Recommended)
1. Go to: https://faucet.polygon.technology/
2. Select "Polygon Amoy Testnet"
3. Paste your wallet address
4. Complete captcha
5. Click "Submit"
6. Receive 0.5 MATIC (can request every 24 hours)

### Option B: Alchemy Faucet
1. Go to: https://www.alchemy.com/faucets/polygon-amoy
2. Sign up/Login to Alchemy (free)
3. Enter your wallet address
4. Receive test MATIC

### Option C: QuickNode Faucet
1. Go to: https://faucet.quicknode.com/polygon/amoy
2. Connect wallet or paste address
3. Tweet to get extra MATIC
4. Receive test tokens

---

## 🔑 Step 3: Setup Verbwire for Testnet

### A. Create Verbwire Account:
1. Go to: https://www.verbwire.com/
2. Sign up for free account
3. Verify your email

### B. Get API Keys:
1. Login to Verbwire dashboard
2. Go to "API Keys" section
3. Create new API key pair
4. Copy both keys:
   - Public Key (pk_live_...)
   - Secret Key (sk_live_...)

### C. Configure for Testnet:
1. In Verbwire dashboard, go to "Settings"
2. Select "Polygon Amoy Testnet" as default chain
3. Save settings

### D. Get Testnet Credits (IMPORTANT):
- Verbwire provides FREE testnet credits
- Contact support or check dashboard for testnet credit options
- Some plans include unlimited testnet minting

---

## 🛠️ Step 4: Configure Your NFTGenie Project

### Update backend/.env:
```env
# Already configured in your project:
VERBWIRE_API_KEY=sk_live_6594b884-3a28-47db-a5c1-0bf497e15ddc
VERBWIRE_PUBLIC_KEY=pk_live_d86f49fc-5be1-44e7-9ee5-d6dba0014d87
CHAIN=polygonAmoy
```

### Verify Chain Configuration:
Your project is already set to use Polygon Amoy testnet ✅

---

## 🎨 Step 5: Mint Your NFT

### Through Frontend:
1. Open http://localhost:3000
2. Click "Connect Wallet"
3. Select MetaMask
4. Make sure you're on Polygon Amoy network
5. Fill the form:
   ```
   Name: My First Testnet NFT
   Image URL: https://picsum.photos/500
   Description: Testing NFT minting on Polygon Amoy
   ```
6. Click "Mint NFT"
7. Approve transaction in MetaMask (uses test MATIC)

### Through API (Direct):
```bash
curl -X POST http://localhost:8000/api/nfts/mint \
  -H "Content-Type: application/json" \
  -d '{
    "name": "API Test NFT",
    "description": "Minted via API on testnet",
    "image_url": "https://via.placeholder.com/500",
    "creator": "YOUR_WALLET_ADDRESS_HERE"
  }'
```

---

## 🔍 Step 6: Verify Your NFT

### Check Transaction:
1. Copy transaction hash from response
2. Go to: https://amoy.polygonscan.com/
3. Paste transaction hash
4. View your minted NFT

### View on OpenSea Testnet:
1. Go to: https://testnets.opensea.io/
2. Connect your wallet
3. Go to your profile
4. See your testnet NFTs

---

## 🆓 FREE Alternatives (No Verbwire Required)

If Verbwire doesn't work, you can use these free alternatives:

### Option 1: ThirdWeb (Recommended for Testing)
```javascript
// Install: npm install @thirdweb-dev/sdk
import { ThirdwebSDK } from "@thirdweb-dev/sdk";

const sdk = new ThirdwebSDK("polygon-amoy");
// Free minting on testnet
```

### Option 2: Direct Smart Contract (Advanced)
1. Deploy your own NFT contract on Polygon Amoy
2. Use Remix IDE: https://remix.ethereum.org/
3. Deploy ERC-721 contract
4. Mint directly to contract

### Option 3: NFT.Storage + Custom Contract
1. Upload images to NFT.Storage (free IPFS)
2. Deploy simple ERC-721 contract
3. Mint with metadata URI

---

## 🚨 Troubleshooting

### Issue: "Insufficient funds"
**Solution**: Get more test MATIC from faucet

### Issue: "API key invalid"
**Solution**: 
- Check if keys are activated in Verbwire dashboard
- Ensure you're using testnet keys
- Contact Verbwire support for testnet access

### Issue: "Network error"
**Solution**:
- Switch MetaMask to Polygon Amoy
- Check RPC URL is working
- Try alternative RPC URLs

### Issue: "Transaction failed"
**Solution**:
- Increase gas fees slightly
- Check wallet has test MATIC
- Verify contract address

---

## 💡 Quick Test Without Real Minting

To test the flow without actual blockchain minting:

1. **Mock Mode**: Update backend to return success without calling Verbwire
2. **Local Blockchain**: Use Ganache for local testing
3. **Hardhat**: Run local Polygon fork

---

## 📞 Getting Help

### Verbwire Support:
- Email: support@verbwire.com
- Discord: Join their Discord server
- Docs: https://docs.verbwire.com/

### Polygon Support:
- Discord: https://discord.gg/polygon
- Forum: https://forum.polygon.technology/

### Free Alternatives:
- ThirdWeb Discord: https://discord.gg/thirdweb
- OpenZeppelin: https://forum.openzeppelin.com/

---

## ✅ Success Indicators

You'll know it's working when:
1. MetaMask shows transaction confirmation
2. Backend returns transaction hash
3. Transaction visible on Polygonscan
4. NFT appears in your wallet
5. Green success message in frontend

---

## 🎯 Next Steps After Success

1. View NFT on OpenSea testnet
2. Test transfer functionality
3. Test marketplace listing
4. Try batch minting
5. Experiment with metadata

---

## 📝 Important Notes

- **Testnet is FREE**: All operations use test tokens
- **Not Real Money**: Test MATIC has no value
- **Reset Anytime**: Can always get new test tokens
- **Practice Safe**: Test everything on testnet first
- **Keep Keys Safe**: Even testnet keys should be protected
