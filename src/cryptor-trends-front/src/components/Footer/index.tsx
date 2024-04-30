import { Link } from "@nextui-org/react";
import { DropdownLanguage } from "../DropdownLanguage";

export function Footer() {
  return (
    <footer className="border-t-1 border-gray-700  py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-y-6">
          <div className="col-span-2 md:col-span-1 md:mx-2">
            <h3 className="text-lg font-semibold text-white">About</h3>
            <p className="mt-2 text-gray-400 text-justify">Welcome to Crypto Trends, your gateway to exploring the dynamic world of cryptocurrencies. Our platform harnesses the power of cutting-edge artificial intelligence to provide insights into the latest trends shaping the crypto landscape. However, it&apos;s important to note that while our platform offers valuable analysis, we do not provide financial advice. As you navigate through Crypto Trends, keep in mind the volatile nature of cryptocurrency markets and the importance of conducting your own research before making any investment decisions.</p>
          </div>
          <div>
            <h3 className="text-lg font-semibold text-white">Services</h3>
            <ul className="mt-2 text-gray-400">
            <li className="mt-1" ><Link href={process.env.NEXT_PUBLIC_API_URL}>API</Link></li>
            <li className="mt-1" ><Link href={process.env.NEXT_PUBLIC_TELEGRAM_URL}>BOT Telegram</Link></li>
            </ul>
          </div>
          <div>
            <h3 className="text-lg font-semibold text-white">Links</h3>
            <ul className="mt-2 text-gray-400">
              <li className="mt-1"><Link href="/#trends">Trends</Link></li>
              <li className="mt-1"><Link href="/about">About</Link></li>
              <li className="mt-1"><Link href="/contact">Contact</Link></li>
            </ul>
          </div>
          <div >
          <h3 className="text-lg font-semibold text-white">Language</h3>
            <DropdownLanguage />
          </div>
        </div>
      </div>
      <div className="mt-8 border-t border-gray-700 pt-4">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <p className="text-center text-gray-400">© {new Date().getFullYear()} Crypto Trends. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
}
