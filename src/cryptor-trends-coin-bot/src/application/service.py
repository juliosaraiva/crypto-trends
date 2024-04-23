from typing import List

from domain.entity import CryptorCoin
from domain.interfaces.service import CryptorCoinService
from domain.interfaces.repository import CryptorCoinRepository
from domain.interfaces.event import CryptorCoinEvent
from libs.coinmarketcap import CoinMarketCap


class CryptorCoinServiceImpl(CryptorCoinService):
    def __init__(
            self,
            cryptor_coin_repository: CryptorCoinRepository,
            cryptor_coin_event: CryptorCoinEvent):
        self.cryptor_coin_repository = cryptor_coin_repository
        self.cryptor_coin_event = cryptor_coin_event

    def get_all_coins(self) -> List[CryptorCoin]:
        self.cryptor_coin_repository.get_all()

    def add_coin(self, collection_name: str, queue_name: str) -> None:
        coinmarketcap = CoinMarketCap()

        # Coin data from CoinMarketCap
        coin_map = coinmarketcap.map()[0]
        coin_historical = coinmarketcap.ohlcv_historical(coin_map["id"])
        coin_quotes = coinmarketcap.quotes_latest(coin_map["id"])
        coin = coinmarketcap.format_coin(coin_map, coin_historical, coin_quotes)

        if coin["historical"] is None:
            print("No historical data found")
            return

        coin_entity = {
            "coin_id": coin["id"],
            "name": coin["name"],
            "symbol": coin["symbol"],
            "rank": coin["rank"],
            "max_supply": coin["max_supply"],
            "circulating_supply": coin["circulating_supply"],
            "total_supply": coin["total_supply"],
            "price": coin["price"],
            "volume_24h": coin["volume_24h"],
            "volume_change_24h": coin["volume_change_24h"],
            "historical": coin["historical"]
        }

        self.cryptor_coin_event.publish_message(queue_name, coin_entity)
        self.cryptor_coin_repository.insert_one(collection_name, coin_entity)
