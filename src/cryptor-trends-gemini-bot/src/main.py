from infrastructure.mongodb_client import MongoDBClient
from infrastructure.rabbitmq_client import RabbitMQClient
from application.service import CryptorCoinIAServiceImpl
from config.settings import settings


# MongoDB connection "mongodb://mongodb:27017/"
mongo_client = MongoDBClient(
  settings.MONGO_HOST,
  settings.MONGO_PORT,
  settings.MONGO_USERNAME,
  settings.MONGO_PASSWORD,
  settings.MONGO_DATABASE
)
mongo_client.connect()

# RabbitMQ connection and create queue
rabbitmq_client = RabbitMQClient(
  settings.RABBITMQ_HOST,
  settings.RABBITMQ_PORT,
  settings.RABBITMQ_USER,
  settings.RABBITMQ_PASSWORD
)
rabbitmq_client.connect()
rabbitmq_client.declare_queue(settings.RABBITMQ_QUEUE_NAME)

# CryptorCoinIAService instance
cryptor_coin_ia_service = CryptorCoinIAServiceImpl(mongo_client, rabbitmq_client)
cryptor_coin_ia_service.add_coin(settings.MONGO_DATABASE, settings.RABBITMQ_QUEUE_NAME)

print("GET")
cryptor_coin_ia_service.get_all_coins()
