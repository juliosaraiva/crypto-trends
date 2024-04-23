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
    Please return JSON describing the trend from this bitcoin historical using the following schema:

    {
      "trend": "high"
    }

    My field is required to be filled out. Please fill it out and return it to me.

    Important: trend is enum with values: low, sideway, high. Only return a single piece of valid JSON text.
    From the next message I will start sending the historical data.

    Example:
    input:
    {
      "name": "Bitcoin",
      "symbol": "BTC",
      "historical":
      [
        {
          "time_open": "2024-04-21T11:00:00.000Z",
          "time_close": "2024-04-21T11:59:59.999Z",
          "time_high": "2024-04-21T11:19:00.000Z",
          "time_low": "2024-04-21T11:01:00.000Z",
          "open": 64980.3767748907,
          "high": 65436.24462599747,
          "low": 64973.487256501605,
          "close": 65269.81132717494,
          "volume": 21997407809.39,
          "market_cap": 1285040734835.02,
          "timestamp": "2024-04-21T11:59:59.999Z"
        }
      ]
    }

    output:
    {
      "trend": "high"
    }
    """
  )

settings = Settings()
