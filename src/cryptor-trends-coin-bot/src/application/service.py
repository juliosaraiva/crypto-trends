from domain.entity import CryptorCoin, CryptorCoinHistorical
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

    def get_all_coins(self):
        self.cryptor_coin_repository.get_all()

    def add_coin(self, collection_name, queue_name):
        coinmarketcap = CoinMarketCap()
        
        # Coin data from CoinMarketCap
        coin_map = coinmarketcap.map()[0]
        coin_historical = coinmarketcap.ohlcv_historical(coin_map["id"])
        coin_quotes = coinmarketcap.quotes_latest(coin_map["id"])
        coin = coinmarketcap.format_coin(coin_map, coin_historical, coin_quotes)

        # Coin data from CoinEntity
        coin_historical_entity = []
        if len(coin['historical']) > 1:
            for value in coin['historical']:
                historical_entity = {
                    "time_open": value["time_open"],
                    "time_close": value["time_close"],
                    "time_high": value["time_high"],
                    "time_low": value["time_low"],
                    "open": value["open"],
                    "high": value["high"],
                    "low": value["low"],
                    "close": value["close"],
                    "volume": value["volume"],
                    "market_cap": value["market_cap"],
                    "timestamp": value["timestamp"]
                }
                coin_historical_entity.append(historical_entity)
        else:
            print("No historical data found")
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
            "historical": coin_historical_entity
        }

        self.cryptor_coin_repository.insert_one(collection_name, coin_entity)
        # self.cryptor_coin_event.publish_message(queue_name, coin)
