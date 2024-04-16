import google.generativeai as genai

# from coinmarketcap import cryptor_trends_dataset
from config.generativeai import generation_config, prompt_parts
from config import envs


data = [
  ("balão verde", "verde"),
  ("caneta preta", "preto"),
]

genai.configure(api_key=envs.GENERATIVEAI_API_KEY)
model = genai.GenerativeModel(model_name="gemini-1.0-pro",
                              generation_config=generation_config)
model.tu(data)
model.save("balao_caneta.pb")

# Carregue o modelo treinado
loaded_model = genai.GenerativeModel.from_file("balao_caneta.pb")
generated_text = loaded_model.generate_content(prompt="Qual a cor do balão?")

# Imprima o texto gerado
print(generated_text)