import os
import textwrap


class Settings:
  MONGO_USERNAME = os.getenv("MONGO_USERNAME")
  MONGO_PASSWORD = os.getenv("MONGO_PASSWORD")
  MONGO_HOST = os.getenv("MONGO_HOST")
  MONGO_PORT = os.getenv("MONGO_PORT")
  MONGO_DATABASE = os.getenv("MONGO_DATABASE")
  MONGO_URL = "mongodb://${MONGO_HOST}:${MONGO_PORT}/"

  RABBITMQ_USER = os.getenv("RABBITMQ_USER")
  RABBITMQ_PASSWORD = os.getenv("RABBITMQ_PASSWORD")
  RABBITMQ_HOST = os.getenv("RABBITMQ_HOST")
  RABBITMQ_PORT = os.getenv("RABBITMQ_PORT")
  RABBITMQ_QUEUE_NAME = os.getenv("RABBITMQ_QUEUE_NAME")
  RABBITMQ_DLQ_QUEUE_NAME = os.getenv("RABBITMQ_DLQ_QUEUE_NAME")
  RABBITMQ_CONSUME_QUEUE_NAME = os.getenv("RABBITMQ_CONSUME_QUEUE_NAME")
  RABBITMQ_CONSUME_DLQ_QUEUE_NAME = os.getenv("RABBITMQ_CONSUME_DLQ_QUEUE_NAME")
  RABBITMQ_CONSUME_MAX_RETRIES = os.getenv("RABBITMQ_CONSUME_MAX_RETRIES") or 5

  GEMINI_API_KEY = os.getenv("GEMINI_API_KEY")
  GEMINI_MODEL_NAME = os.getenv("GEMINI_MODEL_NAME")
  GEMINI_HEADER = textwrap.dedent(
    """
    You are a technical analyst, who have to predict the short-term trend for the cryptocurrencies with symbol "BTC,ETH" based on a technical analysis.
    The dataset below is the historical data by hourly of the bitcoin prices with time_close, open, close, high, low, and volume.

    You have to return the a json with a key trend and a value found for the trend based on your technical analysis:
    - I want you to add a new column to this dataset named trend. Then, based on your technical analysis, add the value related to the trend.
    Must be only one of the following: high, neutral or low

    for each row of the given dataset, you have to add the short-term trend you find using one of the words mentioned before.

    Your output must be only the modified dataset with the new column trend and each row must contains the trend based on your analysis.

    In the trends column, you must return a clear answer with only one of the following words based on your analisis:
    high if it's uptrend
    sideway if it's sideway trend
    low if it's downtrend\n
    You must return the below JSON describing the trend for the cryptocurrency analyzed using historical data:

    {
      "trend": "high"
    }

    Below are a few examples of the historical data for different cryptocurrencies you must analyze and the expected output:
    input:
    {
      "name": "Bitcoin",
      "symbol": "BTC",
      "historical": {
        "time_open": "2024-05-01T14:00:00.000Z",
        "time_close": "2024-05-01T14:59:59.999Z",
        "time_high": "2024-05-01T14:59:00.000Z",
        "time_low": "2024-05-01T14:14:00.000Z",
        "open": 57254.25524180802,
        "high": 57507.73772416699,
        "low": 57080.18096002104,
        "close": 57483.77453818257,
        "volume": 44804101739.47,
        "market_cap": 1132050740351.14,
        "timestamp": "2024-05-01T14:59:59.999Z"
      }
    }

    output:
    {
      "trend": "low"
    }

    input:
    {
      "name": "Ethereum",
      "symbol": "ETH",
      "historical": [
        {
          "time_open": "2024-05-01T14:00:00.000Z",
          "time_close": "2024-05-01T14:59:59.999Z",
          "time_high": "2024-05-01T14:59:00.000Z",
          "time_low": "2024-05-01T14:17:00.000Z",
          "open": 2888.0069184522513,
          "high": 2907.961770644908,
          "low": 2882.1198239080463,
          "close": 2906.2573886550517,
          "volume": 18882009265.35,
          "market_cap": 354752580697.78,
          "timestamp": "2024-05-01T14:59:59.999Z"
        }
    }
    output:
    {
      "trend": "low"
    }

    input:
    {
      "name": "Solana",
      "symbol": "SOL",
      "historical": {
        "time_open": "2024-05-01T14:00:00.000Z",
        "time_close": "2024-05-01T14:59:59.999Z",
        "time_high": "2024-05-01T14:04:00.000Z",
        "time_low": "2024-05-01T14:28:00.000Z",
        "open": 123.37593491226853,
        "high": 123.82835769051542,
        "low": 122.26297880640588,
        "close": 123.22636777447923,
        "volume": 3521769523.57,
        "market_cap": 55154060499.85,
        "timestamp": "2024-05-01T14:59:59.999Z"
      }
    }
    output:
    {
      "trend": "low"
    }
    """
  )

settings = Settings()
