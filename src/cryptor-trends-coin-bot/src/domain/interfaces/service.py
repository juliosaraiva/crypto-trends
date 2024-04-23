from abc import ABC, abstractmethod
from typing import List

from domain.entity import CryptorCoin


class CryptorCoinService(ABC):
    @abstractmethod
    def get_all_coins(self) -> List[CryptorCoin]:
        pass

    @abstractmethod
    def add_coin(self, collection_name: str, queue_name: str) -> None:
        pass