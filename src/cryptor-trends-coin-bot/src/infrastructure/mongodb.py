import pymongo
from domain.entity import CryptorCoin
from domain.repository import CryptorCoinRepository


class MongoCryptorCoinRepository(CryptorCoinRepository):
    def __init__(self, db):
        self.collection = db['coins']

    def save(self, coin: CryptorCoin):
        coin_data = {
            'coin_id': coin.coin_id,
            'name': coin.name,
            'symbol': coin.symbol,
            'rank': coin.rank,
            'max_supply': coin.max_supply,
            'circulating_supply': coin.circulating_supply,
            'total_supply': coin.total_supply,
            'price': coin.price,
            'volume_24h': coin.volume_24h,
            'volume_change_24h': coin.volume_change_24h,
            'time_open': coin.time_open,
            'time_close': coin.time_close,
            'time_high': coin.time_high,
            'time_low': coin.time_low,
            'open': coin.open,
            'high': coin.high,
            'low': coin.low,
            'close': coin.close,
            'volume': coin.volume,
            'market_cap': coin.market_cap,
            'timestamp': coin.timestamp
        }
        self.collection.insert_one(coin_data)

    def get_all(self):
        coins = []
        for coin_data in self.collection.find():
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