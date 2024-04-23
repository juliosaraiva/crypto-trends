import textwrap
import google.generativeai as genai

from domain.interfaces.ia import IA


class GeminiClient(IA):
    def __init__(self):
        self.model = None
        self.config = None
        self.chat = None
        
    def connect(self, api_key: str) -> None:
        genai.configure(api_key=api_key)
        print("Connecting to Gemini AI...")

    def create_model(self, model_name: str) -> None:
        self.model = genai.GenerativeModel(
            model_name=model_name
        )
        self.config = genai.GenerationConfig(
            temperature=0,
            max_output_tokens=2048,
        )

    def start_prompt(self, header: str) -> None:
        self.chat = self.model.start_chat()
        self.chat.send_message(
            textwrap.dedent(header),
            generation_config=self.config
        )

    def generate_content(self, story: str) -> str:
        response = self.chat.send_message(
            textwrap.dedent(story),
            generation_config=self.config
        )
        return response.text