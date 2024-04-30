"use client"
import Loading from "@/app/loading";
import {
  Button,
  Chip,
  ChipProps,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  Input,
  Pagination,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/react";
import { CaretDown, MagnifyingGlass, Ranking } from "@phosphor-icons/react";
import { useAsyncList } from "@react-stately/data";
import { useState } from "react";

const statusColorMap: Record<string, ChipProps["color"]> = {
  high: "success",
  low: "danger",
  neutral: "warning",
};

export function TableList() {
  const [isLoading, setIsLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState(1);
  const [searchValue, setSearchValue] = useState("");
  const itemsPerPage = 10;

  let list = useAsyncList({
    async load({ signal }) {
      try {
        setIsLoading(true);
        const res = await fetch("/api", { signal });
        const json = await res.json();
        setIsLoading(false);
        return { items: json };
      } catch (error) {
        console.error("Failed to fetch data:", error);
        setIsLoading(false);
        return { items: [] };
      }
    },
    async sort({ items, sortDescriptor }: { items: any[]; sortDescriptor: any }) {
      return {
        items: items.sort((a: any, b: any) => {
          let first = a[sortDescriptor.column];
          let second = b[sortDescriptor.column];
          let cmp = (parseInt(first) || first) < (parseInt(second) || second) ? -1 : 1;

          if (sortDescriptor.direction === "descending") {
            cmp *= -1;
          }

          return cmp;
        }),
      };
    },
  });

  const filteredItems = list.items.filter((item: any) =>
    item.name.toLowerCase().includes(searchValue.toLowerCase())
  );

  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  const paginatedItems = filteredItems.slice(startIndex, endIndex);

  const totalPages = Math.ceil(filteredItems.length / itemsPerPage);

  return (
    <div className="max-w-7xl flex flex-col items-center w-full">
      <div className="flex justify-between items-center w-full my-4" id="trends">
        <div>
        <Input
            isClearable
            className="w-full "
            placeholder="Search by name..."
            startContent={<MagnifyingGlass />}
            value={searchValue}
            onClear={() => setSearchValue("")}
            onChange={(e) => setSearchValue(e.target.value)}
          />
        </div>
        <div>
          <Dropdown>
            <DropdownTrigger className="flex">
              <Button variant="flat" endContent={<CaretDown size={14} weight="bold" />}>
                Trends
              </Button>
            </DropdownTrigger>
            <DropdownMenu
              disallowEmptySelection
              aria-label="Table Columns"
              closeOnSelect={false}
              selectionMode="multiple"
            >
              {list.items.map((coin: any) => (
                <DropdownItem key={coin.coin_id} className="capitalize">
                  {coin.trend}
                </DropdownItem>
              ))}
            </DropdownMenu>
          </Dropdown>
        </div>
      </div>

      <Table
        aria-label="Cryptocurrency Table"
        sortDescriptor={list.sortDescriptor}
        onSortChange={list.sort}
        classNames={{
          table: "min-h-[400px]",
        }}
      >
        <TableHeader>
          <TableColumn key="rank" allowsSorting>
            <Ranking size={24} weight="light" />
            Rank
          </TableColumn>
          <TableColumn key="name" allowsSorting>
            Name
          </TableColumn>
          <TableColumn key="circulating_supply" allowsSorting>
            Circulating Supply
          </TableColumn>
          <TableColumn key="price" allowsSorting>
            Price
          </TableColumn>
          <TableColumn key="trend" allowsSorting>
            Trend
          </TableColumn>
        </TableHeader>
        <TableBody isLoading={isLoading} loadingContent={<Loading />}>
          {paginatedItems.map((item: any) => (
            <TableRow key={item.coin_id}>
              <TableCell className="text-yellow-500">{item.rank}</TableCell>
              <TableCell>{item.name}</TableCell>
              <TableCell>${item.circulating_supply}</TableCell>
              <TableCell>${item.price}</TableCell>
              <TableCell>
                <Chip className="capitalize" color={statusColorMap[item.trend]} size="sm" variant="flat">
                  {item.trend}
                </Chip>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <Pagination
        showControls
        showShadow
        total={totalPages}
        page={currentPage}
        onChange={setCurrentPage}
        className="my-4"
      />
    </div>
  );
}
