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
        payload = {
            "id": coin_id,
            "count": 1,
            "interval":"2h",
            "time_period":"hourly",
        }
        response = requests.get(self.url_ohlcv_historical, params=payload, headers=settings.HEADERS)
        coin_historical_dict = json.loads(response.text)
        coin_historical_list = {
            'time_open': coin_historical_dict['data']['quotes'][0]['time_open'],
            'time_close': coin_historical_dict['data']['quotes'][0]['time_close'],
            'time_high': coin_historical_dict['data']['quotes'][0]['time_high'],
            'time_low': coin_historical_dict['data']['quotes'][0]['time_low'],
            'open': coin_historical_dict['data']['quotes'][0]['quote']['USD']['open'],
            'high': coin_historical_dict['data']['quotes'][0]['quote']['USD']['high'],
            'low': coin_historical_dict['data']['quotes'][0]['quote']['USD']['low'],
            'close': coin_historical_dict['data']['quotes'][0]['quote']['USD']['close'],
            'volume': coin_historical_dict['data']['quotes'][0]['quote']['USD']['volume'],
            'market_cap': coin_historical_dict['data']['quotes'][0]['quote']['USD']['market_cap'],
            'timestamp': coin_historical_dict['data']['quotes'][0]['quote']['USD']['timestamp']
        }
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
    