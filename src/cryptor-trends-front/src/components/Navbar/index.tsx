import { Image, Link, Navbar, NavbarBrand, NavbarContent, NavbarItem, NavbarMenu, NavbarMenuItem, NavbarMenuToggle } from "@nextui-org/react";

export default function NavbarHeader() {
  const menuItems = [
    "Trends",
    "About",
    "Contacts",
  ];

  return (
    <Navbar disableAnimation isBordered>
      <NavbarContent className="sm:hidden" justify="start">
        <NavbarMenuToggle />
      </NavbarContent>

      <NavbarContent className="sm:hidden pr-3" justify="center">
        <NavbarBrand>
          <Image src='/images/logo-cryptotrends.svg' alt="" width={50} height={50} />
          <p className="font-bold text-inherit">Crypto Trends</p>
        </NavbarBrand>
      </NavbarContent>

      <NavbarContent className="hidden sm:flex gap-4" justify="center">
        <NavbarBrand>
        <Image src='/images/logo-cryptotrends.svg' alt="" width={50} height={50} />
          <p className="font-bold text-inherit">Crypto Trends</p>
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
          <Link color="foreground" href="#">
            Contact
          </Link>
        </NavbarItem>
      </NavbarContent>


      <NavbarMenu>
        {menuItems.map((item, index) => (
          <NavbarMenuItem key={`${item}-${index}`}>
            <Link
              className="w-full"
              color={
                index === 2 ? "warning" : index === menuItems.length - 1 ? "danger" : "foreground"
              }
              href="#"
              size="lg"
            >
              {item}
            </Link>
          </NavbarMenuItem>
        ))}
      </NavbarMenu>
    </Navbar>
  );
}
