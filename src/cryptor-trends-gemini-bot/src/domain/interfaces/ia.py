from abc import ABC, abstractmethod


class IA(ABC):
    @abstractmethod
    def connect(self, api_key: str) -> None:
        pass

    @abstractmethod
    def create_model(self, model_name: str) -> None:
        pass

    @abstractmethod
    def start_prompt(self, header: str) -> None:
        pass

    @abstractmethod
    def generate_content(self, story: str) -> str:
        pass
    