from pymongo import MongoClient
from typing import List

from domain.interfaces.repository import CryptorCoinRepository
from domain.entity import CryptorCoin


class MongoDBClient(CryptorCoinRepository):
    def __init__(self, host: str, port: str,
                username: str, password: str, database: str):
        self.host = host
        self.port = port
        self.username = username
        self.password = password
        self.database = database
        self.client = None
        self.db = None

    def connect(self) -> None:
        self.client = MongoClient(f"mongodb://{self.host}:{self.port}/")
        self.db = self.client[self.database]
        print(f"Connected to MongoDB server {self.host}:{self.port}")

    def find_all(self, collection: str) -> List[CryptorCoin]:
        coins = []
        for coin_data in self.db[collection].find():
            coin = CryptorCoin(
                coin_id=coin_data['coin_id'],
                name=coin_data['name'],
                symbol=coin_data['symbol'],
                rank=coin_data['rank'],
                max_supply=coin_data['max_supply'],
                circulating_supply=coin_data['circulating_supply'],
                total_supply=coin_data['total_supply'],
                price=coin_data['price'],
                volume_24h=coin_data['volume_24h'],
                volume_change_24h=coin_data['volume_change_24h'],
                time_open=coin_data['time_open'],
                time_close=coin_data['time_close'],
                time_high=coin_data['time_high'],
                time_low=coin_data['time_low'],
                open=coin_data['open'],
                high=coin_data['high'],
                low=coin_data['low'],
                close=coin_data['close'],
                volume=coin_data['volume'],
                market_cap=coin_data['market_cap'],
                timestamp=coin_data['timestamp']
            )
            coins.append(coin)
        return coins

    def insert_one(self, collection: str, data: dict) -> None:
      self.db[collection].insert_one(data)
      print(f"Inserted data into {collection}")

    def close(self) -> None:
        if self.client:
            self.client.close()
