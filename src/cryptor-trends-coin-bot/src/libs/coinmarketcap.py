import requests
import json

from datetime import datetime, timedelta

from config.settings import settings



class CoinMarketCap:
    def __init__(self) -> None:
        self.url_map = settings.COINMARKETCAP_BASE_URL+"/v1/cryptocurrency/map"
        self.url_ohlcv_historical = settings.COINMARKETCAP_BASE_URL+"/v2/cryptocurrency/ohlcv/historical"
        self.url_quotes_latest = settings.COINMARKETCAP_BASE_URL+"/v2/cryptocurrency/quotes/latest"

    def format_coin(self, coin_map, coin_historical, coin_quotes):
        coin = coin_map
        coin['max_supply'] = coin_quotes['max_supply']
        coin['circulating_supply'] = coin_quotes['circulating_supply']
        coin['total_supply'] = coin_quotes['total_supply']
        coin['price'] = coin_quotes['price']
        coin['volume_24h'] = coin_quotes['volume_24h']
        coin['volume_change_24h'] = coin_quotes['volume_change_24h']
        coin['historical'] = coin_historical

        return coin

    def map(self):
        payload = {
            "sort": "cmc_rank",
            "limit": 1,
        }
        response = requests.get(self.url_map, params=payload, headers=settings.HEADERS)
        coin_dict = json.loads(response.text)
        coin_list = []
        for coins in coin_dict['data']:
            row = {
                'id': coins['id'],
                'name': coins['name'],
                'symbol': coins['symbol'],
                'rank': coins['rank']
            }
            coin_list.append(row)
        return coin_list

    def ohlcv_historical(self, coin_id: int):
        time_end = int(datetime.now().timestamp())
        time_start = int((datetime.now() - timedelta(minutes=220)).timestamp())
        payload = {
            "id": [coin_id],
            "time_start":time_start,
            "time_end": time_end,
            "interval":"1h",
            "time_period":"hourly",
        }
        response = requests.get(self.url_ohlcv_historical, params=payload, headers=settings.HEADERS)
        coin_historical_dict = json.loads(response.text)
        coin_historical_list = []
        for quote in coin_historical_dict['data']['quotes']:
            row = {
                'time_open': quote['time_open'],
                'time_close': quote['time_close'],
                'time_high': quote['time_high'],
                'time_low': quote['time_low'],
                'open': quote['quote']['USD']['open'],
                'high': quote['quote']['USD']['high'],
                'low': quote['quote']['USD']['low'],
                'close': quote['quote']['USD']['close'],
                'volume': quote['quote']['USD']['volume'],
                'market_cap': quote['quote']['USD']['market_cap'],
                'timestamp': quote['quote']['USD']['timestamp']
            }
            coin_historical_list.append(row)
        return coin_historical_list
    
    def quotes_latest(self, coin_id: int):
        payload = {
            "id": [coin_id],
            "convert": "USD"
        }
        response = requests.get(self.url_quotes_latest, params=payload, headers=settings.HEADERS)
        coin_dict = json.loads(response.text)
        coin_formated = {
            'max_supply': coin_dict['data'][str(coin_id)]['max_supply'],
            'circulating_supply': coin_dict['data'][str(coin_id)]['circulating_supply'],
            'total_supply': coin_dict['data'][str(coin_id)]['total_supply'],
            'price': coin_dict['data'][str(coin_id)]['quote']['USD']['price'],
            'volume_24h': coin_dict['data'][str(coin_id)]['quote']['USD']['volume_24h'],
            'volume_change_24h': coin_dict['data'][str(coin_id)]['quote']['USD']['volume_change_24h']
        }
        return coin_formated
    