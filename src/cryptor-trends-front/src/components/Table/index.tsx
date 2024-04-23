"use client"
import { Chip, ChipProps, Spinner, Table, TableBody, TableCell, TableColumn, TableHeader, TableRow, } from "@nextui-org/react";
import { useAsyncList } from "@react-stately/data";
import React from "react";
const statusColorMap: Record<string, ChipProps["color"]>  = {
  Stable: "success",
  Growing: "danger",
  // vacation: "warning",
};

export function TableList() {
  const [isLoading, setIsLoading] = React.useState(true);

  let list = useAsyncList({
    async load({signal}) {
      try {
        setIsLoading(true);
        const res = await fetch('/api', { signal });
        const json = await res.json();
        setIsLoading(false);
        return { items: json };
      } catch (error) {
        console.error('Failed to fetch data:', error);
        setIsLoading(false);
        return { items: [] };
      }

    },
    async sort({items, sortDescriptor}: {items: any[], sortDescriptor: any}) {
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

  return (
    <Table
      aria-label="Cryptocurrency Table"
      sortDescriptor={list.sortDescriptor}
      onSortChange={list.sort}
      classNames={{
        table: "min-h-[400px]",
      }}
    >
      <TableHeader>
        <TableColumn key="name" allowsSorting>
          Name
        </TableColumn>
        <TableColumn key="symbol" allowsSorting>
          Symbol
        </TableColumn>
        <TableColumn key="rank" allowsSorting>
          Rank
        </TableColumn>
        <TableColumn key="max_supply" allowsSorting>
          Max Supply
        </TableColumn>
        <TableColumn key="circulating_supply" allowsSorting>
          Circulating Supply
        </TableColumn>
        <TableColumn key="total_supply" allowsSorting>
          Total Supply
        </TableColumn>
        <TableColumn key="price" allowsSorting>
          Price
        </TableColumn>
        <TableColumn key="timestamp" allowsSorting>
          Timestamp
        </TableColumn>
        <TableColumn key="trend" allowsSorting>
          Trend
        </TableColumn>
      </TableHeader>
      <TableBody
        items={list.items}
        isLoading={isLoading}
        loadingContent={<Spinner label="Loading..." />}
      >
        {(item: any) => (
          <TableRow key={item.coin_id}>
            <TableCell>{item.name}</TableCell>
            <TableCell>{item.symbol}</TableCell>
            <TableCell>{item.rank}</TableCell>
            <TableCell>{item.max_supply}</TableCell>
            <TableCell>{item.circulating_supply}</TableCell>
            <TableCell>{item.total_supply}</TableCell>
            <TableCell>{item.price}</TableCell>
            <TableCell>{new Date(item.timestamp).toLocaleString()}</TableCell>
            <TableCell>{}
            <Chip className="capitalize" color={statusColorMap[item.trend]} size="sm" variant="flat">
            {item.trend}
          </Chip>
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
