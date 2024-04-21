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
    """
  )

settings = Settings()
