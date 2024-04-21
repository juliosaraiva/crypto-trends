import time

from infrastructure.mongodb_client import MongoDBClient
from infrastructure.rabbitmq_client import RabbitMQClient
from infrastructure.gemini_client import GeminiClient
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
rabbitmq_client.declare_queue(settings.RABBITMQ_DLQ_QUEUE_NAME)
rabbitmq_client.declare_queue(settings.RABBITMQ_CONSUME_QUEUE_NAME)
rabbitmq_client.declare_queue(settings.RABBITMQ_CONSUME_DLQ_QUEUE_NAME)

# Gemini connection
gemini = GeminiClient()
gemini.connect(settings.GEMINI_API_KEY)
gemini.create_model(settings.GEMINI_MODEL_NAME)
gemini.start_prompt(settings.GEMINI_HEADER)

# CryptorCoinIAService instance
cryptor_coin_ia_service = CryptorCoinIAServiceImpl(mongo_client, rabbitmq_client, gemini)

# RabbitMQ consume message
def callback(ch, method, properties, body):
  retry_count = properties.headers.get('retry_count', 0)
  delivery_tag = method.delivery_tag
  max_retries = settings.RABBITMQ_CONSUME_MAX_RETRIES

  try:
    cryptor_coin_ia_service.add_coin(
      settings.MONGO_DATABASE,
      settings.RABBITMQ_QUEUE_NAME,
      body.decode())
    rabbitmq_client.ack_message(delivery_tag)
  except Exception as e:
    print(f"Error: {e}")
    rabbitmq_client.ack_message(delivery_tag)
    if retry_count < max_retries:
      rabbitmq_client.publish_message(
        queue_name=settings.RABBITMQ_CONSUME_QUEUE_NAME, 
        message=body, 
        retry_count=retry_count + 1)
      time.sleep(3)
    else:
      print(f"Max retries reached: {max_retries}")
      rabbitmq_client.publish_message_to_dlq(
        queue_name=settings.RABBITMQ_CONSUME_DLQ_QUEUE_NAME,
        message=body,
        err=str(e),
        retry_count=retry_count)

rabbitmq_client.add_callback(callback)
rabbitmq_client.consume_message(settings.RABBITMQ_CONSUME_QUEUE_NAME)
