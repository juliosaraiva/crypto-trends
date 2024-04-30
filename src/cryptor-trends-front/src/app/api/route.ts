
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
      trend: "Neutral"
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
      trend: "High"
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
      trend: "Low"
    },
  ];

  return new Response(JSON.stringify(cryptoList), {
    status: 200,
  });
}
