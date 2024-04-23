import { Footer } from "@/components/Footer";
import NavbarHeader from "@/components/Navbar";
import type { Metadata } from "next";
import "./../../public/css/globals.css";
import { Providers } from "./providers";

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
      <body className={'dark text-foreground bg-background '}>
        <Providers>
        <NavbarHeader />
        {children}
        <Footer />
        </Providers>
        </body>
    </html>
  );
}
