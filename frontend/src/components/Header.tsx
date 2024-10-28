import Image from "next/image";
import React from "react";

const Header = () => {
  return (
    <header className="fixed top-0 left-0 right-0 bg-[#383cbc] shadow-md z-50">
      <div className="w-full px-4 sm:px-6 lg:px-8">
        {" "}
        <div className="flex items-center justify-between h-20">
          <div className="flex-shrink-0 pl-0">
            {" "}
            <Image
              className="h-12 w-auto"
              src="https://quest.puffer.fi/images/logos/puffer-logo-white.svg"
              alt="Logo"
            />
          </div>

          <nav className="absolute left-1/2 transform -translate-x-1/2">
            <div className="flex space-x-4">
              <a
                href="#"
                className="px-3 py-2 rounded-md text-xl font-bold text-white hover:bg-gray-800 transition-colors"
              >
                Analytics
              </a>
            </div>
          </nav>

          <div className="flex-shrink-0 w-12"></div>
        </div>
      </div>
    </header>
  );
};

export default Header;
