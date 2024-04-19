from typing import List
from abc import ABC, abstractmethod

from domain.entity import CryptorCoin


class CryptorCoinRepository(ABC):
    @abstractmethod
    def connect(self) -> None:
        pass

    @abstractmethod
    def find_all(self, collection: str) -> List[CryptorCoin]:
        pass

    @abstractmethod
    def insert_one(self, collection: str, data: dict) -> None:
        pass

    @abstractmethod
    def close(self) -> None:
        pass
