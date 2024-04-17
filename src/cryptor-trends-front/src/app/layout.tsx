import { Footer } from "@/components/Footer";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./../../public/css/globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Crypto Trends",
  description: "Crypto Trends description",
  icons: ["/images/favicon.ico"],
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        {children}
        <Footer />
        </body>
    </html>
  );
}
