import ConnectWallet from "@/components/ConnectWallet";
import MintForm from "@/components/MintForm";
import Logo from "@/components/Logo";
import Settings from "@/components/Settings";
import NotificationButton from "@/components/NotificationButton";

export default function Home() {
  return (
    <div className="min-h-screen p-6 sm:p-8 lg:p-12">
      {/* Header */}
      <header className="glass rounded-2xl p-6 mb-12 hover-lift">
        <div className="flex items-center justify-between gap-4">
          <div className="flex items-center gap-3">
            <Logo size={40} />
            <div>
              <h1 className="text-2xl font-bold gradient-text">NFTGenie</h1>
              <p className="text-sm text-slate-400">AI-Powered NFT Platform</p>
            </div>
          </div>
          <div className="flex items-center gap-4">
            <NotificationButton />
            <Settings />
            <ConnectWallet />
          </div>
        </div>
      </header>

      {/* Hero Section */}
      <section className="text-center mb-16">
        <h2 className="text-4xl lg:text-6xl font-bold mb-6">
          Create <span className="gradient-text">Magical</span> NFTs
        </h2>
        <p className="text-xl text-slate-300 max-w-2xl mx-auto mb-12">
          Harness the power of AI to create, mint, and discover unique digital assets on Polygon network
        </p>
      </section>

      {/* Main Content */}
      <main className="grid grid-cols-1 lg:grid-cols-2 gap-12 max-w-7xl mx-auto">
        {/* Mint Form */}
        <section className="space-y-8">
          <MintForm />
        </section>

        {/* Info Cards */}
        <section className="space-y-8">
          {/* Feature Tags */}
          <div className="text-center">
            <div className="flex flex-wrap justify-center gap-4 text-sm">
              <span className="glass px-4 py-2 rounded-full">🚀 Fast Minting</span>
              <span className="glass px-4 py-2 rounded-full">🤖 AI-Powered</span>
              <span className="glass px-4 py-2 rounded-full">💎 Low Fees</span>
              <span className="glass px-4 py-2 rounded-full">🔒 Secure</span>
            </div>
          </div>

          {/* How it Works */}
          <div className="glass rounded-2xl p-6 hover-lift">
            <div className="flex items-center gap-3 mb-4">
              <div className="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-500 to-green-500 flex items-center justify-center">
                <span className="text-sm">📋</span>
              </div>
              <h3 className="text-xl font-semibold">How it Works</h3>
            </div>
            <ol className="space-y-3 text-slate-300">
              <li className="flex items-start gap-3">
                <span className="w-6 h-6 rounded-full bg-blue-500/20 text-blue-400 text-sm flex items-center justify-center mt-0.5 font-semibold">1</span>
                <span>Connect your wallet using the button above</span>
              </li>
              <li className="flex items-start gap-3">
                <span className="w-6 h-6 rounded-full bg-green-500/20 text-green-400 text-sm flex items-center justify-center mt-0.5 font-semibold">2</span>
                <span>Fill out the mint form with your NFT details</span>
              </li>
              <li className="flex items-start gap-3">
                <span className="w-6 h-6 rounded-full bg-amber-500/20 text-amber-400 text-sm flex items-center justify-center mt-0.5 font-semibold">3</span>
                <span>Our backend processes the mint on Polygon Amoy</span>
              </li>
              <li className="flex items-start gap-3">
                <span className="w-6 h-6 rounded-full bg-green-500/20 text-green-400 text-sm flex items-center justify-center mt-0.5 font-semibold">4</span>
                <span>View your transaction on Polygonscan</span>
              </li>
            </ol>
          </div>

          {/* Features */}
          <div className="glass rounded-2xl p-6 hover-lift">
            <div className="flex items-center gap-3 mb-4">
              <div className="w-8 h-8 rounded-lg bg-gradient-to-br from-green-500 to-blue-500 flex items-center justify-center">
                <span className="text-sm">✨</span>
              </div>
              <h3 className="text-xl font-semibold">Features</h3>
            </div>
            <div className="grid grid-cols-2 gap-4">
              <div className="text-center p-3 rounded-xl bg-blue-500/10 border border-blue-500/20">
                <div className="text-2xl mb-2">⚡</div>
                <div className="text-sm font-medium">Lightning Fast</div>
              </div>
              <div className="text-center p-3 rounded-xl bg-green-500/10 border border-green-500/20">
                <div className="text-2xl mb-2">🔐</div>
                <div className="text-sm font-medium">Secure</div>
              </div>
              <div className="text-center p-3 rounded-xl bg-amber-500/10 border border-amber-500/20">
                <div className="text-2xl mb-2">💰</div>
                <div className="text-sm font-medium">Low Cost</div>
              </div>
              <div className="text-center p-3 rounded-xl bg-green-500/10 border border-green-500/20">
                <div className="text-2xl mb-2">🌍</div>
                <div className="text-sm font-medium">Eco-Friendly</div>
              </div>
            </div>
          </div>
        </section>
      </main>
      

    </div>
  );
}
