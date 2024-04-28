import { HeroSection } from "@/components/HeroSection";
import { TableList } from "@/components/Table";

export function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-4 sm:p-16 bg-no-repeat bg-[url('/images/looper-pattern.svg')]">
      <HeroSection />
      <TableList />




    </main>
  );
}
