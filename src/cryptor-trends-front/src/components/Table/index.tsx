"use client"
import Loading from '@/app/loading';
import { Coin } from '@/types/coin';
import { ColorMap } from '@/types/statusColorMap';
import { capitalize } from '@/utils/capitalize';
import { formatFloat } from '@/utils/formatFloat';
import { removeDuplicates } from '@/utils/removeDuplicates';
import { Button, Chip, Dropdown, DropdownItem, DropdownMenu, DropdownTrigger, Input, Pagination, Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@nextui-org/react';
import { CaretDown, MagnifyingGlass, Ranking } from '@phosphor-icons/react';
import { useAsyncList } from '@react-stately/data';
import { useEffect, useMemo, useState } from 'react';
const statusColorMap: ColorMap = {
  high: "success",
  low: "danger",
  sideway: "warning",
};

export function TableList() {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [searchValue, setSearchValue] = useState<string>("");
  const [selectedTrends, setSelectedTrends] = useState<string[]>([]);
  const itemsPerPage: number = 10;

  let list = useAsyncList<Coin>({
    async load({ signal }) {
      try {
        setIsLoading(true);
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL }`, { signal });
        const json = await res.json();
        setIsLoading(false);
        if (json === null || json === undefined) {
          console.error("Failed to fetch data: Unable to retrieve data");
          return { items: [] };
        }

        return { items: json };
      } catch (error) {
        console.error("Failed to fetch data:", error);
        setIsLoading(false);
        return { items: [] };
      }
    },
    async sort({ items, sortDescriptor }: { items: Coin[]; sortDescriptor: any }) {
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

  useEffect(() => {
    // Select all trends initially
    const allTrends = removeDuplicates(list.items.map((coin: Coin) => coin.trend));
    setSelectedTrends(allTrends);
  }, [list.items]);

  const filteredItems = useMemo(() => {
    return list.items.filter((item: Coin) =>
      item.name.toLowerCase().includes(searchValue.toLowerCase()) &&
      selectedTrends.includes(item.trend) // Filter based on selected trends
    );
  }, [list.items, searchValue, selectedTrends]);

  const startIndex: number = (currentPage - 1) * itemsPerPage;
  const endIndex: number = startIndex + itemsPerPage;
  const paginatedItems: Coin[] = filteredItems.slice(startIndex, endIndex);

  const totalPages: number = Math.ceil(filteredItems.length / itemsPerPage);

  const handleToggleTrend = (trend: string) => {
    if (selectedTrends.includes(trend)) {
      setSelectedTrends(selectedTrends.filter(item => item !== trend));
    } else {
      setSelectedTrends([...selectedTrends, trend]);
    }
  };

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
              selectedKeys={selectedTrends}
            >
              {removeDuplicates(list.items.map((coin: Coin) => coin.trend)).map((uniqueTrend: string, index: number) => (
                <DropdownItem
                  key={uniqueTrend}
                  className="capitalize"
                  onClick={() => handleToggleTrend(uniqueTrend)}
                >
                  {capitalize(uniqueTrend)}
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
          {paginatedItems.map((item: Coin) => (
            <TableRow key={item.coin_id}>
              <TableCell className="text-yellow-500">{item.rank}</TableCell>
              <TableCell>{item.name}</TableCell>
              <TableCell>${formatFloat(item.circulating_supply)}</TableCell>
              <TableCell>${formatFloat(item.price)}</TableCell>
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
