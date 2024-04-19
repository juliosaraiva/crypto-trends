from abc import ABC, abstractclassmethod
from typing import List

from domain.entity import CryptorCoinIA


class CryptorCoinIAService(ABC):
  @abstractclassmethod
  def get_all_coins(self) -> List[CryptorCoinIA]:
    pass

  @abstractclassmethod
  def add_coin(self, collection_name: str, queue_name: str) -> None:
    pass
