from typing import List
from abc import ABC, abstractclassmethod

from domain.entity import CryptorCoinIA


class CryptorCoinIARepository(ABC):
  @abstractclassmethod
  def connect(self) -> None:
    pass

  @abstractclassmethod
  def find_all(self, collection: str) -> list[CryptorCoinIA]:
    pass

  @abstractclassmethod
  def insert_one(self, collection: str, data: dict) -> None:
    pass

  def close(self) -> None:
    pass
