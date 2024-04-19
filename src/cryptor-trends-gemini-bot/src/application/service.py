from domain.interfaces.service import CryptorCoinIAService
from domain.interfaces.repository import CryptorCoinIARepository
from domain.interfaces.event import CryptorCoinIAEvent
from domain.entity import CryptorCoinIA


class CryptorCoinIAServiceImpl(CryptorCoinIAService):
  def __init__(self,
              cryptor_coin_ia_repository: CryptorCoinIARepository,
              cryptor_coin_ia_event: CryptorCoinIAEvent):
    self.cryptor_coin_ia_repository = cryptor_coin_ia_repository
    self.cryptor_coin_ia_event = cryptor_coin_ia_event

  def get_all_coins(self) -> List[CryptorCoinIA]:
    pass

  def add_coin(self, collection_name: str, queue_name: str) -> None:
    pass
    # self.cryptor_coin_ia_event.publish_message(queue_name=queue_name, message={"message": "ok"})
    # self.cryptor_coin_ia_repository.insert_one(collection=collection_name, data={"message": "ok"})
