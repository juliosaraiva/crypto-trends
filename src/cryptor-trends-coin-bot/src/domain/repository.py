from abc import ABC, abstractmethod
from typing import List
from domain.entity import CryptorCoin


class CryptorCoinRepository(ABC):
    @abstractmethod
    def get_all(self) -> List[CryptorCoin]:
        pass

    @abstractmethod
    def save(self, coin: CryptorCoin) -> None:
        pass
    