import requests
from pymongo import MongoClient
from domain.entity import CryptorCoin
from infrastructure.mongodb import MongoCryptorCoinRepository
from infrastructure.service import CryptorCoinServiceImpl
from coinmarketcap import coinmarketcap_map, coinmarketcap_ohlcv_historical

from config.settings import settings


client = MongoClient("mongodb://mongodb:27017/")
db = client[settings.MONGO_DATABASE]

cryptor_coin_repository = MongoCryptorCoinRepository(db)
cryptor_coin_service = CryptorCoinServiceImpl(cryptor_coin_repository)

coins = coinmarketcap_map()

for coin in coins:
    coins_history = coinmarketcap_ohlcv_historical(coin)
    coin["historical_data"] = coins_history

print(coins)

# coin = CryptorCoin(
#     coin_id=1,
#     name="Bitcoin",
#     symbol="BTC",
#     rank=1,
#     max_supply=21000000,
#     circulating_supply=18700000,
#     total_supply=21000000,
#     price=50000,
#     volume_24h=1000000000,
#     volume_change_24h=0.05,
#     time_open=10000,
#     time_close=20000,
#     time_high=30000,
#     time_low=40000,
#     open=45000,
#     high=60000,
#     low=40000,
#     close=50000,
#     volume=1000000,
#     market_cap=1000000000,
#     timestamp=1630000000
# )

# cryptor_coin_service.save_coin(coin)
