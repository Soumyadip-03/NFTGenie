import Image from "next/image";

import ConnectWallet from "@/components/ConnectWallet";
import MintForm from "@/components/MintForm";

export default function Home() {
  return (
    <div className="font-sans min-h-screen p-8 sm:p-12">
      <header className="flex items-center justify-between gap-4 mb-10">
        <h1 className="text-2xl font-bold">NFTGenie • Polygon Amoy</h1>
        <ConnectWallet />
      </header>

      <main className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <section>
          <MintForm />
        </section>
        <section className="bg-black/10 rounded-lg p-4">
          <h2 className="text-xl font-semibold mb-2">How it works</h2>
          <ol className="list-decimal list-inside space-y-2 text-sm opacity-90">
            <li>Connect your wallet using the button above.</li>
            <li>Fill out the Quick Mint form and submit.</li>
            <li>Backend calls Verbwire Quick Mint on Polygon Amoy.</li>
            <li>View the transaction on Polygonscan.</li>
          </ol>
        </section>
      </main>
    </div>
  );
}
