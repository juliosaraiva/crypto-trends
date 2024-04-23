from typing import List
from abc import ABC, abstractmethod

from domain.entity import CryptorCoinIA


class CryptorCoinIARepository(ABC):
  @abstractmethod
  def connect(self) -> None:
    pass

  @abstractmethod
  def find_all(self, collection: str) -> list[CryptorCoinIA]:
    pass

  @abstractmethod
  def insert_one(self, collection: str, data: dict) -> None:
    pass

  def close(self) -> None:
    pass
