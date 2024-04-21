from abc import ABC, abstractmethod
from typing import List

from domain.entity import CryptorCoinIA


class CryptorCoinIAService(ABC):
  @abstractmethod
  def get_all_coins(self) -> List[CryptorCoinIA]:
    pass

  @abstractmethod
  def add_coin(self, collection_name: str, queue_name: str, data: str) -> None:
    pass
