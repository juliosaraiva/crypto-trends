from abc import ABC, abstractmethod


class CryptorCoinIAEvent(ABC):
  @abstractmethod
  def connect(self) -> None:
    pass

  @abstractmethod
  def add_callback(self, callback) -> None:
    pass

  @abstractmethod
  def declare_queue(self, queue_name: str) -> None:
    pass

  @abstractmethod
  def publish_message(self, queue_name: str, message: dict, retry_count=0) -> None:
    pass

  @abstractmethod
  def publish_message_to_dlq(self, queue_name: str, message: dict, err: str, retry_count=0) -> None:
    pass

  @abstractmethod
  def consume_message(self, queue_name: str) -> str:
    pass

  @abstractmethod
  def ack_message(self, delivery_tag: int) -> None:
    pass

  @abstractmethod
  def nack_message(self, delivery_tag: int) -> None:
    pass

  @abstractmethod
  def close(self) -> None:
    pass
