"use client";

import Header from "../components/Header";
import PufferChart from "@/components/PufferChart";
import { useState, useEffect } from "react";

export default function Home() {
  const [showTitles, setShowTitles] = useState(false);

  useEffect(() => {
    setTimeout(() => {
      setShowTitles(true);
    }, 100);
  }, []);

  return (
    <>
      <div
        className="fixed inset-0 w-full h-full -z-10 bg-cover bg-center bg-no-repeat opacity-100"
        style={{
          backgroundImage: 'url("/background.png")',
        }}
      />

      <Header />

      <div className="min-h-screen pt-20 p-8 pb-20 gap-16 sm:p-20 sm:pt-24 font-[family-name:var(--font-geist-sans)]">
        <div className="w-full max-w-7xl mx-auto">
          <div className="flex flex-col items-center mb-6">
            <div
              style={{
                transform: showTitles ? "translateY(0)" : "translateY(20px)",
                opacity: showTitles ? 1 : 0,
                transition: "all 600ms cubic-bezier(0.34, 1.56, 0.64, 1)",
              }}
            >
              <h1 className="text-white text-5xl md:text-6xl lg:text-6xl font-bold tracking-wider font-serif mb-4">
                Puffer Analytics
              </h1>
            </div>

            <div
              style={{
                transform: showTitles ? "translateY(0)" : "translateY(-20px)",
                opacity: showTitles ? 1 : 0,
                transition: "all 600ms cubic-bezier(0.34, 1.56, 0.64, 1)",
              }}
            >
              <h4 className="text-white text-xl md:text-2xl opacity-80 text-center">
                Analysing the Future of Ethereum
              </h4>
            </div>
          </div>

          {/* Static chart */}
          <PufferChart />
        </div>
      </div>
    </>
  );
}
