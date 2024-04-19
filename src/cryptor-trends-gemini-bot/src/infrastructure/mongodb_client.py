from pymongo import MongoClient

from domain.interfaces.repository import CryptorCoinIARepository
from domain.entity import CryptorCoinIA


class MongoDBClient(CryptorCoinIARepository):
  def __init__(self, host, port, userame, password, database):
    self.host = host
    self.port = port
    self.username = username
    self.password = password
    self.database = database
    self.client = None
    self.db = None

  def connect(self) -> None:
    self.client = MongoClient(f"mongodb://{self.host}:{self.port}")
    self.db = self.client[self.database]
    print(f"Connected to MongoDB server {self.host}:{self.port}")

  def find_all(self, collection: str) -> list[CryptorCoinIA]:
    coins = []
    for coin_data in self.db[collection].find_all():
      coin = CryptorCoinIA(
        coin_id=coin_data["coin_data"],
        name=coin_data["name"],
        symbol=coin_data["symbol"],
        rank=coin_data["rank"],
        price=coin_data["price"],
        status=coin_data["status"]
      )
      coins.append(coin)
    return coins

  def insert_one(self, collection: str, data: dict) -> None:
    self.db[collection].insert_one(data)
    print(f"Inserted data into {collection}")

  def close(self) -> None:
    if self.client:
      self.client.close()
