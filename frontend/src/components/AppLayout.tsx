import Image from "next/image";
import React from "react";

const AppLayout = () => {
  return (
    <div className="min-h-screen relative">
      {/* Background Image - Replace /path-to-your-image.png with your actual image path */}
      <div
        className="fixed inset-0 w-full h-full -z-10 bg-cover bg-center bg-no-repeat"
        style={{
          backgroundImage: 'url("/public/background.png")',
          opacity: 0.15, // Adjust this value to make background lighter/darker
        }}
      />

      <header className="w-full bg-white shadow-md">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center h-16">
            <div className="flex-shrink-0">
              <Image
                className="h-8 w-8"
                src="/public/puffer-logo.png"
                alt="Company Logo"
              />
            </div>

            {/* Navigation */}
            <nav className="ml-6">
              <div className="flex space-x-4">
                <a
                  href="#"
                  className="px-3 py-2 rounded-md text-sm font-medium text-gray-900 hover:bg-gray-100"
                >
                  Rate
                </a>
              </div>
            </nav>
          </div>
        </div>
      </header>

      {/* Main Content Area */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        {/* Your content will go here */}
      </main>
    </div>
  );
};

export default AppLayout;
