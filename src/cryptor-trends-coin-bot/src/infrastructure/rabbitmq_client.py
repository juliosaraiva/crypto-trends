import pika
import json

from domain.interfaces.event import CryptorCoinEvent

class RabbitMQClient(CryptorCoinEvent):
    def __init__(self, host: str, port: str,
                username: str, password: str) -> None:
        self.host = host
        self.port = port
        self.username = username
        self.password = password
        self.connection = None
        self.channel = None

    def connect(self) -> None:
        credentials = pika.PlainCredentials(self.username, self.password)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host=self.host, port=self.port, credentials=credentials)
        )
        self.channel = self.connection.channel()
        print('Connected to RabbitMQ')

    def declare_queue(self, queue_name: str) -> None:
        self.channel.queue_declare(queue=queue_name)

    def publish_message(self, queue_name: str, message: dict) -> None:
        self.channel.basic_publish(exchange='', routing_key=queue_name, body=json.dumps(message))
        print(f"Published message to {queue_name}")

    def close(self) -> None:
        if self.connection:
            self.connection.close()
