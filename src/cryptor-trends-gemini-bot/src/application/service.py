import json

from typing import List

from domain.interfaces.service import CryptorCoinIAService
from domain.interfaces.repository import CryptorCoinIARepository
from domain.interfaces.event import CryptorCoinIAEvent
from domain.interfaces.ia import IA
from domain.entity import CryptorCoinIA


class CryptorCoinIAServiceImpl(CryptorCoinIAService):
  def __init__(self,
              cryptor_coin_ia_repository: CryptorCoinIARepository,
              cryptor_coin_ia_event: CryptorCoinIAEvent,
              cryptor_coin_ia: IA):
    self.cryptor_coin_ia_repository = cryptor_coin_ia_repository
    self.cryptor_coin_ia_event = cryptor_coin_ia_event
    self.cryptor_coin_ia = cryptor_coin_ia

  def get_all_coins(self) -> List[CryptorCoinIA]:
    return self.cryptor_coin_ia_repository.find_all()

  def add_coin(self, collection_name: str, queue_name: str, data: str) -> None:
    resp = self.cryptor_coin_ia.generate_content(data)
    resp_json = json.loads(resp)

    data_json = json.loads(data)
    data_json["trend"] = resp_json["trend"]

    coin={
      "coin_id": data_json["coin_id"],
      "name": data_json["name"],
      "symbol": data_json["symbol"],
      "rank": data_json["rank"],
      "max_supply": data_json["max_supply"],
      "circulating_supply": data_json["circulating_supply"],
      "total_supply": data_json["total_supply"],
      "price": data_json["price"],
      "timestamp": data_json["historical"]["timestamp"],
      "trend": data_json["trend"]
    }
    self.cryptor_coin_ia_event.publish_message(
      queue_name=queue_name,
      message=json.dumps(coin))
    self.cryptor_coin_ia_repository.insert_one(
      collection=collection_name,
      data=coin)
