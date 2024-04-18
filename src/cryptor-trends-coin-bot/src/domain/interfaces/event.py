from abc import ABC, abstractmethod


class CryptorCoinEvent(ABC):
    @abstractmethod
    def connect(self) -> None:
        pass

    @abstractmethod
    def declare_queue(self, queue_name: str) -> None:
        pass

    @abstractmethod
    def publish_message(self, queue_name: str, message: str) -> None:
        pass

    @abstractmethod
    def close(self) -> None:
        pass
    