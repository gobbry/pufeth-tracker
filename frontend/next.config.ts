import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  experimental: {
    optimizePackageImports: ["@chakra-ui/react"],
  },
  output: "export",
  images: {
    unoptimized: true,
  },
};

export default nextConfig;
