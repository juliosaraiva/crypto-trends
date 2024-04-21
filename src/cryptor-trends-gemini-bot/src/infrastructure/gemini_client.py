import textwrap
import google.generativeai as genai

from domain.interfaces.ia import IA


class GeminiClient(IA):
    def __init__(self):
        self.model = None
        
    def connect(self, api_key: str) -> None:
        genai.configure(api_key=api_key)
        print("Connecting to Gemini AI...")

    def create_model(self, model_name: str) -> None:
        self.model = genai.GenerativeModel(
            model_name=model_name
        )

    def start_prompt(self, header: str) -> None:
        self.model.generate_content(
            textwrap.dedent(header),
            generation_config={'response_mime_type':'application/json'}
        )

    def generate_content(self, story: str) -> str:
        response = self.model.generate_content(
            textwrap.dedent(story),
            generation_config={'response_mime_type':'application/json'}
        )
        return response.text