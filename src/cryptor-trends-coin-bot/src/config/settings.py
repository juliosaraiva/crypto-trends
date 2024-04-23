import os


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

    COINMARKETCAP_API_KEY = os.getenv("COINMARKETCAP_API_KEY")
    COINMARKETCAP_BASE_URL = "https://pro-api.coinmarketcap.com"

    HEADERS = {
        'Content-type': 'application/json',
        'Accept-Encoding': 'deflate,gzip',
        'X-CMC_PRO_API_KEY': os.getenv("COINMARKETCAP_API_KEY")
    }

settings = Settings()
