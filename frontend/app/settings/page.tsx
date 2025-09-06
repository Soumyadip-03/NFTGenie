"use client";

import { useState } from "react";
import { useTheme } from "@/contexts/ThemeContext";
import Logo from "@/components/Logo";
import Link from "next/link";

export default function SettingsPage() {
  const { theme, setTheme } = useTheme();
  const [notifications, setNotifications] = useState(true);
  const [autoConnect, setAutoConnect] = useState(false);
  const [language, setLanguage] = useState("en");

  const handleReset = () => {
    setTheme("dark");
    setNotifications(true);
    setAutoConnect(false);
    setLanguage("en");
  };

  return (
    <div className="min-h-screen p-6 sm:p-8 lg:p-12">
      {/* Header */}
      <header className="glass rounded-2xl p-6 mb-8 hover-lift">
        <div className="flex items-center justify-between gap-4">
          <Link href="/" className="flex items-center gap-3 hover:opacity-80 transition-opacity">
            <Logo size={40} />
            <div>
              <h1 className="text-2xl font-bold gradient-text">NFTGenie</h1>
              <p className="text-sm text-slate-400">AI-Powered NFT Platform</p>
            </div>
          </Link>
          <Link 
            href="/"
            className="glass rounded-xl p-3 hover-lift transition-all duration-300"
          >
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
              <path d="M19 12H5M12 19L5 12L12 5" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
            </svg>
          </Link>
        </div>
      </header>

      {/* Settings Content */}
      <div className="max-w-2xl mx-auto">
        <div className="glass rounded-2xl p-8 hover-lift">
          <div className="text-center mb-8">
            <h2 className="text-3xl font-bold gradient-text mb-2">Settings</h2>
            <p className="text-slate-400">Customize your NFTGenie experience</p>
          </div>

          <div className="space-y-8">
            {/* Theme Selection */}
            <div>
              <label className="text-lg font-semibold text-slate-300 mb-4 block">🎨 Theme</label>
              <div className="grid grid-cols-2 gap-4">
                {["dark", "light"].map((t) => (
                  <button
                    key={t}
                    onClick={() => setTheme(t)}
                    className={`p-4 rounded-xl border transition-all ${
                      theme === t
                        ? "border-purple-500/50 bg-purple-500/10 text-purple-400"
                        : "border-white/20 hover:border-white/30 bg-slate-800/50"
                    }`}
                  >
                    <div className="text-2xl mb-2">{t === "dark" ? "🌙" : "☀️"}</div>
                    <div className="font-medium">{t.charAt(0).toUpperCase() + t.slice(1)} Mode</div>
                  </button>
                ))}
              </div>
            </div>

            {/* Language Selection */}
            <div>
              <label className="text-lg font-semibold text-slate-300 mb-4 block">🌍 Language</label>
              <select
                value={language}
                onChange={(e) => setLanguage(e.target.value)}
                className="w-full p-4 rounded-xl glass border border-white/20 focus:border-cyan-500/50 focus:outline-none focus:ring-2 focus:ring-cyan-500/20 transition-all text-lg font-bold dark:text-slate-100 dark:bg-slate-900/95 light:bg-white"
                style={{ color: document.documentElement.classList.contains('light') ? '#111827' : '#f1f5f9' }}
              >
                <option value="en" className="dark:bg-slate-900 light:bg-white dark:text-slate-100 light:text-gray-900 font-bold py-2">🇺🇸 English</option>
                <option value="es" className="dark:bg-slate-900 light:bg-white dark:text-slate-100 light:text-gray-900 font-bold py-2">🇪🇸 Español</option>
                <option value="fr" className="dark:bg-slate-900 light:bg-white dark:text-slate-100 light:text-gray-900 font-bold py-2">🇫🇷 Français</option>
                <option value="de" className="dark:bg-slate-900 light:bg-white dark:text-slate-100 light:text-gray-900 font-bold py-2">🇩🇪 Deutsch</option>
                <option value="ja" className="dark:bg-slate-900 light:bg-white dark:text-slate-100 light:text-gray-900 font-bold py-2">🇯🇵 日本語</option>
              </select>
            </div>

            {/* Preferences */}
            <div>
              <label className="text-lg font-semibold text-slate-300 mb-4 block">⚙️ Preferences</label>
              <div className="space-y-6">
                <div className="flex items-center justify-between p-4 rounded-xl bg-slate-800/30 border border-white/10">
                  <div>
                    <div className="text-lg font-medium text-slate-300">🔔 Notifications</div>
                    <div className="text-sm text-slate-400">Get notified about transactions and updates</div>
                  </div>
                  <button
                    onClick={() => setNotifications(!notifications)}
                    className={`relative w-14 h-7 rounded-full transition-colors ${
                      notifications ? "bg-purple-500" : "bg-slate-600"
                    }`}
                  >
                    <div
                      className={`absolute top-0.5 w-6 h-6 bg-white rounded-full transition-transform ${
                        notifications ? "translate-x-7" : "translate-x-0.5"
                      }`}
                    />
                  </button>
                </div>

                <div className="flex items-center justify-between p-4 rounded-xl bg-slate-800/30 border border-white/10">
                  <div>
                    <div className="text-lg font-medium text-slate-300">🔗 Auto Connect</div>
                    <div className="text-sm text-slate-400">Automatically connect your wallet on visit</div>
                  </div>
                  <button
                    onClick={() => setAutoConnect(!autoConnect)}
                    className={`relative w-14 h-7 rounded-full transition-colors ${
                      autoConnect ? "bg-cyan-500" : "bg-slate-600"
                    }`}
                  >
                    <div
                      className={`absolute top-0.5 w-6 h-6 bg-white rounded-full transition-transform ${
                        autoConnect ? "translate-x-7" : "translate-x-0.5"
                      }`}
                    />
                  </button>
                </div>
              </div>
            </div>

            {/* Action Buttons */}
            <div className="flex gap-4 pt-6 border-t border-white/20">
              <button
                onClick={handleReset}
                className="flex-1 p-4 rounded-xl border border-white/20 hover:border-white/30 bg-slate-800/50 hover:bg-slate-800/70 transition-colors font-medium"
              >
                🔄 Reset to Default
              </button>
              <Link
                href="/"
                className="flex-1 p-4 rounded-xl bg-gradient-to-r from-purple-500 to-cyan-500 hover:from-purple-600 hover:to-cyan-600 transition-all font-medium text-center"
              >
                ✅ Save & Return
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}