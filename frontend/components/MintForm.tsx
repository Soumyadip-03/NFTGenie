"use client";

import { useState } from "react";
import { useAccount } from "wagmi";

export default function MintForm() {
  const { address, isConnected } = useAccount();
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
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <form onSubmit={onMint} className="w-full max-w-xl space-y-4 bg-black/10 rounded-lg p-4">
      <h2 className="text-xl font-semibold">Quick Mint (Polygon Amoy)</h2>
      <input
        className="w-full p-2 rounded bg-black/20 border border-white/10"
        placeholder="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
      />
      <input
        className="w-full p-2 rounded bg-black/20 border border-white/10"
        placeholder="Image URL (IPFS or HTTPS)"
        value={imageUrl}
        onChange={(e) => setImageUrl(e.target.value)}
        required
      />
      <textarea
        className="w-full p-2 rounded bg-black/20 border border-white/10"
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        rows={3}
      />
      <button
        className="px-4 py-2 rounded bg-purple-600 hover:bg-purple-700 disabled:opacity-50"
        disabled={loading || !isConnected}
      >
        {loading ? "Minting..." : "Mint NFT"}
      </button>
      {tx && (
        <div className="text-green-400">
          Minted! Tx: <a className="underline" target="_blank" href={`https://amoy.polygonscan.com/tx/${tx}`}>{tx}</a>
        </div>
      )}
      {error && <div className="text-red-400">{error}</div>}
    </form>
  );
}

