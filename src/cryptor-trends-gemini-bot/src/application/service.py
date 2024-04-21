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
    # resp = self.cryptor_coin_ia.generate_content(data)
    # print(resp)
    mock_data={
      "coin_id": 1,
      "name": "Bitcoin",
      "symbol": "BTC",
      "rank": 1,
      "max_supply": 21000000,
      "circulating_supply": 19687987,
      "total_supply": 19687987,
      "price": 65031.34477218155,
      "timestamp": "2024-04-21T02:59:59.999Z",
      "trend": "high"
    }
    self.cryptor_coin_ia_event.publish_message(
      queue_name=queue_name,
      message=json.dumps(mock_data))
    self.cryptor_coin_ia_repository.insert_one(
      collection=collection_name,
      data=mock_data)
