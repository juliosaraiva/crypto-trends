from abc import ABC, abstractclassmethod


class CryptorCoinIAEvent(ABC):
  @abstractclassmethod
  def connect(self) -> None:
    pass

  @abstractclassmethod
  def declare_queue(self, queue_name: str) -> None:
    pass

  @abstractclassmethod
  def publish_message(self, queue_name: str, message: dict) -> None:
    pass

  @abstractclassmethod
  def consume_message(self, queue_name) -> str:
    pass

  @abstractclassmethod
  def close(self) -> None:
    pass
