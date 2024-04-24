package infrastructure

import (
	"fmt"

	"github.com/juliosaraiva/crypto-trends/src/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQClient(config *config.Settings) (*amqp.Connection, error) {
	uri := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.RabbitMQUser,
		config.RabbitMQPassword,
		config.RabbitMQHost,
		config.RabbitMQPort,
	)

	conn, err := amqp.Dial(uri)
	return conn, err
}

func Channel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	return ch, err
}

func QueueDeclare(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(queueName, true, false, false, false, nil)
}

func Close(conn *amqp.Connection) error {
	return conn.Close()
}

func CloseChannel(ch *amqp.Channel) error {
	return ch.Close()
}
