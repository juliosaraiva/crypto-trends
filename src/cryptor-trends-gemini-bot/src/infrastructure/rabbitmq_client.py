import pika

from domain.interfaces.event import CryptorCoinIAEvent


class RabbitMQClient(CryptorCoinIAEvent):
  def __init__(self, host: str, port: str,
              username: str, password: str) -> None:
    self.host = host
    self.port = port
    self.username = username
    self.password = password
    self.callback = None
    self.connection = None
    self.channel = None

  def connect(self) -> None:
    credentials = pika.PlainCredentials(self.username, self.password)
    self.connection = pika.BlockingConnection(
      pika.ConnectionParameters(host=self.host, port=self.port, credentials=credentials)
    )
    self.channel = self.connection.channel()
    print('Connected to RabbitMQ')

  def add_callback(self, callback) -> None:
    self.callback = callback

  def declare_queue(self, queue_name: str) -> None:
    self.channel.queue_declare(queue_name, durable=True)

  def publish_message(self, queue_name: str, message: dict, retry_count=0) -> None:
    self.channel.basic_publish(
      exchange='',
      routing_key=queue_name,
      body=message,
      properties=pika.BasicProperties(
        delivery_mode=2,  # make message persistent
        headers={'retry_count': retry_count}
      ))
    print(f"Published message to {queue_name}")

  def publish_message_to_dlq(self, queue_name: str, message: dict, err: str, retry_count=0) -> None:
    self.channel.basic_publish(
      exchange='',
      routing_key=queue_name,
      body=message,
      properties=pika.BasicProperties(
        delivery_mode=2,  # make message persistent
        headers={'error': err, 'retry_count': retry_count}
      ))
    print(f"Published dql message to {queue_name}")

  def consume_message(self, queue_name: str) -> str:
    print(f"Consuming message from {queue_name}")
    self.channel.basic_consume(
      queue=queue_name,
      on_message_callback=self.callback)
    print(' [*] Waiting for messages. To exit press CTRL+C')
    self.channel.start_consuming()

  def ack_message(self, delivery_tag: int) -> None:
    self.channel.basic_ack(delivery_tag=delivery_tag)

  def nack_message(self, delivery_tag: int) -> None:
    self.channel.basic_nack(delivery_tag=delivery_tag)

  def close(self) -> None:
    if self.connection:
      self.connection.close()

