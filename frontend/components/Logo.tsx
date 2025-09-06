export default function Logo({ size = 40 }: { size?: number }) {
  return (
    <div 
      className="relative flex items-center justify-center rounded-xl bg-gradient-to-br from-purple-600 via-blue-600 to-cyan-600 p-0.5 shadow-lg"
      style={{ width: size, height: size }}
    >
      <div className="w-full h-full bg-gradient-to-br from-gray-950 via-slate-950 to-gray-950 rounded-lg flex items-center justify-center relative overflow-hidden border border-gray-800/30">
        {/* Magic Wand */}
        <svg 
          width={size * 0.6} 
          height={size * 0.6} 
          viewBox="0 0 24 24" 
          fill="none" 
          className="relative z-10"
        >
          {/* Wand stick */}
          <path 
            d="M3 21L21 3" 
            stroke="url(#gradient1)" 
            strokeWidth="2" 
            strokeLinecap="round"
          />
          {/* Star at tip */}
          <path 
            d="M21 3L19 5L21 7L23 5L21 3Z" 
            fill="url(#gradient2)"
          />
          {/* Sparkles */}
          <circle cx="8" cy="16" r="1" fill="url(#gradient3)" />
          <circle cx="12" cy="12" r="1.5" fill="url(#gradient4)" />
          <circle cx="6" cy="10" r="0.8" fill="url(#gradient5)" />
          
          <defs>
            <linearGradient id="gradient1" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#8b5cf6" />
              <stop offset="100%" stopColor="#06b6d4" />
            </linearGradient>
            <linearGradient id="gradient2" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#f59e0b" />
              <stop offset="100%" stopColor="#f97316" />
            </linearGradient>
            <linearGradient id="gradient3" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#ec4899" />
              <stop offset="100%" stopColor="#8b5cf6" />
            </linearGradient>
            <linearGradient id="gradient4" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#06b6d4" />
              <stop offset="100%" stopColor="#3b82f6" />
            </linearGradient>
            <linearGradient id="gradient5" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stopColor="#10b981" />
              <stop offset="100%" stopColor="#06b6d4" />
            </linearGradient>
          </defs>
        </svg>
        
        {/* Animated glow effect */}
        <div className="absolute inset-0 bg-gradient-to-br from-purple-600/10 via-blue-600/10 to-cyan-600/10 animate-pulse rounded-lg"></div>
      </div>
    </div>
  );
}