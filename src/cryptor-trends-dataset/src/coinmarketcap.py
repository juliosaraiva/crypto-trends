import requests
import json

import pandas as pd

from config import envs


def cryptor_trends_dataset():
    url = "https://pro-api.coinmarketcap.com/v2/cryptocurrency/ohlcv/historical"
    payload = {
        "id": [1],
        "convert":"USD"
    }
    headers = {
        'Content-type': 'application/json',
        'Accept-Encoding': 'deflate,gzip',
        'X-CMC_PRO_API_KEY': envs.COINMARKETCAP_API_KEY
    }

    response = requests.get(url, params=payload, headers=headers)

    data_dict = json.loads(response.text)
    data_list = []

    # id, name, symbol, time_open, time_close, time_high, time_low, open, high, low, close, volume, market_cap, timestamp
    for quote in data_dict['data']['quotes']:
        row = {
            'id': data_dict['data']['id'],
            'name': data_dict['data']['name'],
            'symbol': data_dict['data']['symbol'],
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
        data_list.append(row)

    df = pd.DataFrame(data_list)
    csv_string = df.to_csv(index=False)

    return csv_string
