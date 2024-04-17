from domain.service import CryptorCoinService
from domain.repository import CryptorCoinRepository


class CryptorCoinServiceImpl(CryptorCoinService):
    def __init__(self, cryptor_coin_repository: CryptorCoinRepository):
        self.cryptor_coin_repository = cryptor_coin_repository

    def get_all_coins(self):
        self.cryptor_coin_repository.get_all()

    def save_coin(self, coin):
        self.cryptor_coin_repository.save(coin)
