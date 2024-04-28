'use client'
import { Button, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger } from "@nextui-org/react";

export function DropdownLanguage() {
  return (
    <Dropdown>
      <DropdownTrigger>
        <Button
          variant="bordered"
        >
          Select Language
        </Button>
      </DropdownTrigger>
      <DropdownMenu aria-label="Static Actions">
        <DropdownItem key="portuguese">Portuguese</DropdownItem>
        <DropdownItem key="english">English</DropdownItem>
        <DropdownItem key="spanish">Spanish</DropdownItem>
        <DropdownItem key="french">French</DropdownItem>
      </DropdownMenu>
    </Dropdown>
  );
}
