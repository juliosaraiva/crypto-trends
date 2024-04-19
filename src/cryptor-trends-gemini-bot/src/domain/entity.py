class CryptorCoinIA:
  def __init__(self, coin_id, name, symbol, rank, price, status):
    self.coin_id = coin_id
    self.name = name
    self.symbol = symbol
    self.rank = rank
    self.price = price
    self.status = status

    self.validate()

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
    if not self.price:
      empyt_fields.append('price')
    if not self.status:
      empyt_fields.append('status')
    if empyt_fields:
      raise ValueError(f"Empty fields: {'. '.join(empyt_fields)}")
