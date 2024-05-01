"use client"
import { Button, Image, Link } from "@nextui-org/react";
import Lottie from "react-lottie";
import animationData from './crypto_trends.json';
export function HeroSection(){
  const defaultOptions = {
    loop: true,
    autoplay: true,
    animationData: animationData,
    rendererSettings: {
      preserveAspectRatio: "xMidYMid slice",
    },
  };
  return (
    <section>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="py-12 md:py-20 flex flex-col md:flex-row items-center justify-between">
          <div className="w-full md:w-1/2 mb-8 md:mb-0 md:order-1">
            <div className="flex justify-center">
              <Image src="/images/logo-cryptotrends.svg" alt="Hero" className="text-center" width={100} height={100} />
            </div>
            <h2 className="text-3xl font-bold text-primary text-center md:text-left">Welcome to <span className="text-white">Crypto Trends</span></h2>
            <p className="mt-4 text-gray-400 text-center md:text-left">The easiest way to quick identify trends in cryptocurrencies. Our platform harnesses the power of cutting-edge artificial intelligence to provide insights into the latest trends shaping the crypto landscape.</p>
            <div className="flex justify-center md:justify-start mt-6">
              <Link className="mr-4 flex-1" href="/about">
                <Button className="flex-1" color="primary" variant="shadow">
                  Learn More
                </Button>
              </Link>
              <Link className="flex-1" href="#trends">
                <Button className="flex-1" color="primary" variant="ghost" >
                  Trends
                </Button>
              </Link>
            </div>
          </div>
          <div className="w-full md:w-1/2 md:order-2 text-center md:text-right">
            <Lottie options={defaultOptions} style={{ width: '100%', height: '100%' }} />
          </div>
        </div>
      </div>
    </section>
  );
}
