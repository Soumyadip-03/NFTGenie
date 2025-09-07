"use client";

import { useState } from "react";
import { useAccount } from "wagmi";
import Logo from "@/components/Logo";
import { useNotifications } from "@/contexts/NotificationContext";

export default function MintForm() {
  const { address, isConnected } = useAccount();
  const { addNotification } = useNotifications();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [imageUrl, setImageUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [tx, setTx] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function onMint(e: React.FormEvent) {
    e.preventDefault();
    if (!isConnected || !address) {
      setError("Connect your wallet first");
      return;
    }
    setLoading(true);
    setError(null);
    setTx(null);

    // Add "minting started" notification
    addNotification({
      type: 'info',
      title: 'NFT Minting Started',
      message: `Creating "${name}" on Polygon Amoy...`
    });

    try {
      const res = await fetch(`/backend/api/nfts/mint`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name,
          description,
          image_url: imageUrl,
          creator: address,
          chain: "polygonAmoy",
        }),
      });
      const data = await res.json();
      if (!res.ok || data.success === false) {
        throw new Error(data.message || "Mint failed");
      }
      setTx(data.transaction_hash || "");
      
      // Add success notification
      addNotification({
        type: 'success',
        title: 'NFT Minted Successfully! 🎉',
        message: `"${name}" has been created and is now live on the blockchain.`
      });
    } catch (err: any) {
      setError(err.message);
      
      // Add error notification
      addNotification({
        type: 'error',
        title: 'Minting Failed',
        message: `Failed to mint "${name}": ${err.message}`
      });
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="glass rounded-2xl p-6 hover-lift">
      <div className="flex items-center gap-3 mb-6">
        <Logo size={32} />
        <div>
          <h2 className="text-xl font-semibold">Create Your NFT</h2>
          <p className="text-sm text-slate-400">Mint on Polygon Amoy Testnet</p>
        </div>
      </div>

      <form onSubmit={onMint} className="space-y-5">
        <div className="space-y-2">
          <label className="text-sm font-medium text-slate-300">NFT Name *</label>
          <input
            className="w-full p-4 rounded-xl glass border border-white/10 focus:border-blue-500/50 focus:outline-none focus:ring-2 focus:ring-blue-500/20 transition-all placeholder-slate-500"
            placeholder="Enter your NFT name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>

        <div className="space-y-2">
          <label className="text-sm font-medium text-slate-300">Image URL *</label>
          <input
            className="w-full p-4 rounded-xl glass border border-white/10 focus:border-green-500/50 focus:outline-none focus:ring-2 focus:ring-green-500/20 transition-all placeholder-slate-500"
            placeholder="https://... or ipfs://..."
            value={imageUrl}
            onChange={(e) => setImageUrl(e.target.value)}
            required
          />
          <p className="text-xs text-slate-400">Supported: IPFS, HTTPS URLs</p>
        </div>

        <div className="space-y-2">
          <label className="text-sm font-medium text-slate-300">Description</label>
          <textarea
            className="w-full p-4 rounded-xl glass border border-white/10 focus:border-amber-500/50 focus:outline-none focus:ring-2 focus:ring-amber-500/20 transition-all placeholder-slate-500 resize-none"
            placeholder="Describe your NFT..."
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            rows={4}
          />
        </div>

        <button
          className={`w-full p-4 rounded-xl font-semibold transition-all duration-300 ${
            loading || !isConnected
              ? 'bg-slate-600 cursor-not-allowed opacity-50'
              : 'bg-gradient-to-r from-blue-500 to-green-500 hover:from-blue-600 hover:to-green-600 glow-blue hover:glow-green'
          }`}
          disabled={loading || !isConnected}
        >
          {loading ? (
            <div className="flex items-center justify-center gap-2">
              <div className="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
              Minting Your NFT...
            </div>
          ) : !isConnected ? (
            'Connect Wallet to Mint'
          ) : (
            <div className="flex items-center justify-center gap-2">
              <span>✨</span>
              Mint NFT
              <span>✨</span>
            </div>
          )}
        </button>

        {tx && (
          <div className="glass rounded-xl p-4 border border-green-500/30 bg-green-500/10">
            <div className="flex items-center gap-2 mb-2">
              <span className="text-green-400">✓</span>
              <span className="font-semibold text-green-400">Successfully Minted!</span>
            </div>
            <p className="text-sm text-slate-300 mb-3">Your NFT has been created on Polygon Amoy</p>
            <a 
              className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-green-500/20 hover:bg-green-500/30 text-green-400 text-sm font-medium transition-colors" 
              target="_blank" 
              href={`https://amoy.polygonscan.com/tx/${tx}`}
            >
              <span>🔗</span>
              View on Polygonscan
            </a>
          </div>
        )}

        {error && (
          <div className="glass rounded-xl p-4 border border-red-500/30 bg-red-500/10">
            <div className="flex items-center gap-2 mb-2">
              <span className="text-red-400">⚠</span>
              <span className="font-semibold text-red-400">Minting Failed</span>
            </div>
            <p className="text-sm text-slate-300">{error}</p>
          </div>
        )}
      </form>
    </div>
  );
}

