import { Image, Link, Navbar, NavbarBrand, NavbarContent, NavbarItem, NavbarMenu, NavbarMenuToggle } from "@nextui-org/react";

export default function NavbarHeader() {
  const menuItems = [
    "Trends",
    "About",
    "Teams",
  ];

  return (
    <Navbar disableAnimation isBordered>
      <NavbarContent className="sm:hidden" justify="start">
        <NavbarMenuToggle />
      </NavbarContent>

      <NavbarContent className="sm:hidden pr-3" justify="center">
        <NavbarBrand>
          <Link href="/">
            <Image src='/images/logo-cryptotrends.svg' alt="" width={50} height={50} />
            <p className="font-bold text-inherit">Crypto Trends</p>
            </Link>
        </NavbarBrand>
      </NavbarContent>

      <NavbarContent className="hidden sm:flex gap-4" justify="center">
      <NavbarBrand>
          <Link href="/">
            <Image src='/images/logo-cryptotrends.svg' alt="" width={50} height={50} />
            <p className="font-bold text-inherit">Crypto Trends</p>
            </Link>
        </NavbarBrand>
        <NavbarItem>
          <Link color="foreground" href="/#trends">
            Trends
          </Link>
        </NavbarItem>
        <NavbarItem >
          <Link href="/about" color="foreground" aria-current="page" >
            About
          </Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="/team">
            Team
          </Link>
        </NavbarItem>
      </NavbarContent>


      <NavbarMenu>
        <NavbarItem>
            <Link color="foreground" href="/#trends">
              Trends
            </Link>
          </NavbarItem>
        <NavbarItem>
            <Link color="foreground" href="/about">
              About
            </Link>
          </NavbarItem>
        <NavbarItem>
            <Link color="foreground" href="/team">
              Team
            </Link>
          </NavbarItem>
      </NavbarMenu>
    </Navbar>
  );
}
