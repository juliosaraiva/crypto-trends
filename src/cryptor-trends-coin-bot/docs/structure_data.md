# Regra de negocio: sempre pegar a ultima hora.
# vai roda em 1 e 1 hora

# https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest
# {
#     "max_supply": 21000000,
#     "circulating_supply": 18700000,
#     "total_supply": 21000000,
#     "price": 67840.573581488,
#     "volume_24h": 19967785808.97",
#     "volume_change_24h": 0.05
# }

# https://pro-api.coinmarketcap.com/v2/cryptocurrency/ohlcv/historical
# {
#     "time_open": "2024-04-06T00:00:00.000Z",
#     "time_close": "2024-04-06T23:59:59.999Z",
#     "time_high": "2024-04-06T23:10:00.000Z",
#     "time_low": "2024-04-06T02:38:00.000Z",
#     "open": 67840.573581488,
#     "high": 69629.60204854626,
#     "low": 67491.71523520295,
#     "close": 68896.10998427868,
#     "volume": 19967785808.97,
#     "market_cap": 1355566990096.02,
#     "timestamp": "2024-04-06T23:59:59.999Z"
# }


# https://pro-api.coinmarketcap.com/v1/cryptocurrency/map

# {
# 	"id": 1,
# 	"rank": 1,
# 	"name": "Bitcoin",
# 	"symbol": "BTC",
# },

# final
# {
#     "id": 1,
#     "name": "Bitcoin",
#     "symbol": "BTC",
#     "rank": 1,
#     "max_supply": 21000000,
#     "circulating_supply": 18700000,
#     "total_supply": 21000000,
#     "price": 67840.573581488,
#     "volume_24h": 19967785808.97",
#     "volume_change_24h": 0.05
#     "historical_data": {
#         "time_open": "2024-04-06T00:00:00.000Z",
#         "time_close": "2024-04-06T23:59:59.999Z",
#         "time_high": "2024-04-06T23:10:00.000Z",
#         "time_low": "2024-04-06T02:38:00.000Z",
#         "open": 67840.573581488,
#         "high": 69629.60204854626,
#         "low": 67491.71523520295,
#         "close": 68896.10998427868,
#         "volume": 19967785808.97,
#         "market_cap": 1355566990096.02,
#         "timestamp": "2024-04-06T23:59:59.999Z"
#     }
# }