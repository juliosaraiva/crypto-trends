class CryptorCoin:
    def __init__(self, coin_id, name, symbol, rank, 
                 time_open, max_supply, circulating_supply,
                 total_supply, price, volume_24h, volume_change_24h,
                 time_close, time_high, time_low,
                 open, high, low, close, volume,
                 market_cap, timestamp) -> None:
        self.coin_id = coin_id
        self.name = name
        self.symbol = symbol
        self.rank = rank
        self.max_supply = max_supply
        self.circulating_supply = circulating_supply
        self.total_supply = total_supply
        self.price = price
        self.volume_24h = volume_24h
        self.volume_change_24h = volume_change_24h
        self.time_open = time_open
        self.time_close = time_close
        self.time_high = time_high
        self.time_low = time_low
        self.open = open
        self.high = high
        self.low = low
        self.close = close
        self.volume = volume
        self.market_cap = market_cap
        self.timestamp = timestamp

        self.validate()

    def __str__(self) -> str:
        return f"{self.name} ({self.symbol}) - {self.close} USD"
    
    def __eq__(self, value: object) -> bool:
        if not isinstance(value, CryptorCoin):
            return False
        return self.coin_id == value.coin_id
    
    def __hash__(self) -> int:
        return hash(self.coin_id)
    
    def validate(self):
        empyt_fields = []
        if not self.coin_id:
            empyt_fields.append('coin_id')
        if not self.name:
            empyt_fields.append('name')
        if not self.symbol:
            empyt_fields.append('symbol')
        if not self.rank:
            empyt_fields.append('rank')
        if not self.max_supply:
            empyt_fields.append('max_supply')
        if not self.circulating_supply:
            empyt_fields.append('circulating_supply')
        if not self.total_supply:
            empyt_fields.append('total_supply')
        if not self.price:
            empyt_fields.append('price')
        if not self.volume_24h:
            empyt_fields.append('volume_24h')
        if not self.volume_change_24h:
            empyt_fields.append('volume_change_24h')
        if not self.time_open:
            empyt_fields.append('time_open')
        if not self.time_close:
            empyt_fields.append('time_close')
        if not self.time_high:
            empyt_fields.append('time_high')
        if not self.time_low:
            empyt_fields.append('time_low')
        if not self.open:
            empyt_fields.append('open')
        if not self.high:
            empyt_fields.append('high')
        if not self.low:
            empyt_fields.append('low')
        if not self.close:
            empyt_fields.append('close')
        if not self.volume:
            empyt_fields.append('volume')
        if not self.market_cap:
            empyt_fields.append('market_cap')
        if not self.timestamp:
            empyt_fields.append('timestamp')
        if empyt_fields:
            raise ValueError(f"Empty fields: {', '.join(empyt_fields)}")
