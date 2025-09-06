import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Explicitly set the root directory for Turbopack
  experimental: {
    turbo: {
      root: "./",
    },
  },
  async rewrites() {
    return [
      {
        source: "/backend/:path*",
        destination: "http://localhost:8000/:path*",
      },
    ];
  },
};

export default nextConfig;
