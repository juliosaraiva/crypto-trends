import requests
import json
from config.settings import settings
from datetime import datetime, timedelta


def coinmarketcap_map():
    url = settings.COINMARKETCAP_BASE_URL+"/v1/cryptocurrency/map"

    payload = {
        "sort": "cmc_rank",
        "limit": 1,
    }

    response = requests.get(url, params=payload, headers=settings.HEADERS)
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

def coinmarketcap_ohlcv_historical(coin):
    url = settings.COINMARKETCAP_BASE_URL+"/v2/cryptocurrency/ohlcv/historical"

    time_end = int(datetime.now().timestamp())
    time_start = int((datetime.now() - timedelta(minutes=120)).timestamp())

    payload = {
        "id": [coin['id']],
        "time_start":time_start,
        "time_end": time_end,
        "interval":"1h",
        "time_period":"hourly",
    }

    response = requests.get(url, params=payload, headers=settings.HEADERS)

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