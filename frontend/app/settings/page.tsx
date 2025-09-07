"use client";

import { useState } from "react";
import Link from "next/link";

export default function SettingsPage() {
  const [notifications, setNotifications] = useState({
    transactions: true,
    updates: false,
    marketing: false,
  });

  const handleNotificationChange = (type: keyof typeof notifications) => {
    setNotifications(prev => ({
      ...prev,
      [type]: !prev[type]
    }));
  };

  return (
    <div className="min-h-screen p-6 sm:p-8 lg:p-12">
      {/* Header */}
      <header className="glass rounded-2xl p-6 mb-8 hover-lift">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold gradient-text">Settings</h1>
          <Link href="/" className="text-blue-400 hover:text-blue-300 transition-colors">
            ← Back
          </Link>
        </div>
      </header>

      {/* Settings Content */}
      <main className="max-w-2xl mx-auto space-y-6">
        {/* Preferences Section */}
        <section className="glass rounded-2xl p-6 hover-lift">
          <h2 className="text-xl font-semibold mb-6 flex items-center gap-2">
            <span>⚙️</span>
            Preferences
          </h2>

          {/* Notifications */}
          <div className="space-y-4">
            <h3 className="text-lg font-medium flex items-center gap-2">
              <span>🔔</span>
              Notifications
            </h3>
            
            <div className="space-y-3 ml-8">
              <div className="flex items-center justify-between p-3 rounded-xl glass">
                <div>
                  <div className="font-medium text-slate-200">Transaction Updates</div>
                  <div className="text-sm text-slate-400">Get notified when your NFTs are minted or transferred</div>
                </div>
                <button
                  onClick={() => handleNotificationChange('transactions')}
                  className={`w-12 h-6 rounded-full transition-all duration-300 ${
                    notifications.transactions 
                      ? 'bg-green-500' 
                      : 'bg-slate-600'
                  }`}
                >
                  <div className={`w-5 h-5 bg-white rounded-full transition-transform duration-300 ${
                    notifications.transactions ? 'translate-x-6' : 'translate-x-0.5'
                  }`}></div>
                </button>
              </div>

              <div className="flex items-center justify-between p-3 rounded-xl glass">
                <div>
                  <div className="font-medium text-slate-200">Platform Updates</div>
                  <div className="text-sm text-slate-400">New features, improvements, and announcements</div>
                </div>
                <button
                  onClick={() => handleNotificationChange('updates')}
                  className={`w-12 h-6 rounded-full transition-all duration-300 ${
                    notifications.updates 
                      ? 'bg-green-500' 
                      : 'bg-slate-600'
                  }`}
                >
                  <div className={`w-5 h-5 bg-white rounded-full transition-transform duration-300 ${
                    notifications.updates ? 'translate-x-6' : 'translate-x-0.5'
                  }`}></div>
                </button>
              </div>

              <div className="flex items-center justify-between p-3 rounded-xl glass">
                <div>
                  <div className="font-medium text-slate-200">Marketing & Tips</div>
                  <div className="text-sm text-slate-400">NFT trends, tips, and promotional content</div>
                </div>
                <button
                  onClick={() => handleNotificationChange('marketing')}
                  className={`w-12 h-6 rounded-full transition-all duration-300 ${
                    notifications.marketing 
                      ? 'bg-green-500' 
                      : 'bg-slate-600'
                  }`}
                >
                  <div className={`w-5 h-5 bg-white rounded-full transition-transform duration-300 ${
                    notifications.marketing ? 'translate-x-6' : 'translate-x-0.5'
                  }`}></div>
                </button>
              </div>
            </div>
          </div>
        </section>

        {/* Other Settings */}
        <section className="glass rounded-2xl p-6 hover-lift">
          <h2 className="text-xl font-semibold mb-6 flex items-center gap-2">
            <span>🎨</span>
            Appearance
          </h2>
          
          <div className="p-3 rounded-xl glass">
            <div className="font-medium text-slate-200 mb-2">Theme</div>
            <div className="text-sm text-slate-400">Dark mode (Light mode coming soon)</div>
          </div>
        </section>

        {/* Account Settings */}
        <section className="glass rounded-2xl p-6 hover-lift">
          <h2 className="text-xl font-semibold mb-6 flex items-center gap-2">
            <span>👤</span>
            Account
          </h2>
          
          <div className="space-y-3">
            <div className="p-3 rounded-xl glass">
              <div className="font-medium text-slate-200 mb-2">Wallet Connection</div>
              <div className="text-sm text-slate-400">Manage your connected wallets and permissions</div>
            </div>
            
            <div className="p-3 rounded-xl glass">
              <div className="font-medium text-slate-200 mb-2">Privacy</div>
              <div className="text-sm text-slate-400">Control your data and privacy settings</div>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
}