# WalletConnect Setup

To fix the "Connection interrupted while trying to subscribe" error:

## Get a WalletConnect Project ID

1. Go to https://cloud.walletconnect.com/
2. Sign up or log in
3. Create a new project
4. Copy the Project ID
5. Update `.env.local` with your Project ID:

```
NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID=your-actual-project-id-here
```

## Alternative: Disable WalletConnect (Development Only)

For local development, you can temporarily disable WalletConnect by using only injected wallets (MetaMask).

The current setup uses a fallback that should work for basic testing.