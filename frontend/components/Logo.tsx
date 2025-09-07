export default function Logo({ size = 40 }: { size?: number }) {
  return (
    <div 
      className="relative flex items-center justify-center rounded-full bg-gradient-to-br from-blue-500 via-green-500 to-blue-600 p-0.5 shadow-lg"
      style={{ width: size, height: size }}
    >
      <div className="w-full h-full bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 rounded-full flex items-center justify-center relative overflow-hidden border border-slate-700/50">
        {/* Diamond/Gem Icon */}
        <svg 
          width={size * 0.6} 
          height={size * 0.6} 
          viewBox="0 0 24 24" 
          fill="none" 
          className="relative z-10"
        >
          {/* Diamond shape */}
          <path 
            d="M6 3h12l4 6-10 12L2 9l4-6z" 
            fill="url(#gradient1)"
            stroke="url(#gradient2)"
            strokeWidth="1"
          />
          {/* Inner facets */}
          <path 
            d="M6 3l6 6 6-6M2 9l10 12L22 9M12 9v12" 
            stroke="url(#gradient3)"
            strokeWidth="0.8"
            opacity="0.7"
          />
          {/* Sparkles around diamond */}
          <circle cx="4" cy="6" r="0.8" fill="url(#gradient4)" />
          <circle cx="20" cy="6" r="0.8" fill="url(#gradient5)" />
          <circle cx="19" cy="15" r="1" fill="url(#gradient4)" />
          
          <defs>
            <linearGradient id="gradient1" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#3b82f6" />
              <stop offset="50%" stopColor="#10b981" />
              <stop offset="100%" stopColor="#3b82f6" />
            </linearGradient>
            <linearGradient id="gradient2" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#10b981" />
              <stop offset="100%" stopColor="#3b82f6" />
            </linearGradient>
            <linearGradient id="gradient3" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#ffffff" />
              <stop offset="100%" stopColor="#e2e8f0" />
            </linearGradient>
            <linearGradient id="gradient4" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#3b82f6" />
              <stop offset="100%" stopColor="#10b981" />
            </linearGradient>
            <linearGradient id="gradient5" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#10b981" />
              <stop offset="100%" stopColor="#3b82f6" />
            </linearGradient>
          </defs>
        </svg>
        
        {/* Animated glow effect */}
        <div className="absolute inset-0 bg-gradient-to-br from-blue-500/10 via-green-500/10 to-blue-600/10 animate-pulse rounded-full"></div>
      </div>
    </div>
  );
}