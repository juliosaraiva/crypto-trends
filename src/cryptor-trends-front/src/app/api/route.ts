
interface CryptoData {
  coin_id: string;
  name: string;
  symbol: string;
  rank: number;
  max_supply: number;
  circulating_supply: number;
  total_supply: number;
  price: number;
  timestamp: Date;
  trend: string;
}

export async function GET() {
  const cryptoList: CryptoData[] = [
    {
      coin_id: "bitcoin_id",
      name: "Bitcoin",
      symbol: "BTC",
      rank: 1,
      max_supply: 21000000,
      circulating_supply: 18750375,
      total_supply: 18750375,
      price: 45000.00,
      timestamp: new Date("2024-04-22"),
      trend: "sideway"
    },
    {
      coin_id: "ethereum_id",
      name: "Ethereum",
      symbol: "ETH",
      rank: 2,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-22"),
      trend: "high"
    },
    {
      coin_id: "moeda_id",
      name: "Moeda",
      symbol: "MOED",
      rank: 3,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "1",
      name: "Moeda1",
      symbol: "MOED",
      rank: 4,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "2",
      name: "Moeda2",
      symbol: "MOED",
      rank: 5,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "3",
      name: "Moeda3",
      symbol: "MOED",
      rank: 6,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "4",
      name: "Moeda4",
      symbol: "MOED",
      rank: 7,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "5",
      name: "Moeda5",
      symbol: "MOED",
      rank: 8,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "6",
      name: "Moeda6",
      symbol: "MOED",
      rank: 9,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "7",
      name: "Moeda7",
      symbol: "MOED",
      rank: 10,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "8",
      name: "Moeda8",
      symbol: "MOED",
      rank: 11,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
    {
      coin_id: "9",
      name: "Moeda9",
      symbol: "MOED",
      rank: 12,
      max_supply: 116018364,
      circulating_supply: 116018364,
      total_supply: 116018364,
      price: 3000.00,
      timestamp: new Date("2024-04-23"),
      trend: "low"
    },
  ];

  return new Response(JSON.stringify(cryptoList), {
    status: 200,
  });
}
