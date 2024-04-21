package config

import "os"

type Settings struct {
	Domain        string
	Version       string
	WebServerPort string
	Environment   string
	// MongoDB
	MongoDBHost string
	MongoDBPort string
	// MongoDBUser     string
	// MongoDBPassword string
	MongoDBName string
	// RabbitMQ
	RabbitMQHost                string
	RabbitMQPort                string
	RabbitMQUser                string
	RabbitMQPassword            string
	RabbitMQConsumeQueueName    string
	RabbitMQConsumeDLQQueueName string
}

func NewSettings() *Settings {
	return &Settings{
		Domain:        "http://localhost",
		Version:       "v1.0.0",
		WebServerPort: ":8080",
		Environment:   "development",
		// MongoDB
		MongoDBHost: os.Getenv("MONGODB_HOST"),
		MongoDBPort: os.Getenv("MONGODB_PORT"),
		// MongoDBUser:     "",
		// MongoDBPassword: "",
		MongoDBName: os.Getenv("MONGO_DATABASE"),
		// RabbitMQ
		RabbitMQHost:                os.Getenv("RABBITMQ_HOST"),
		RabbitMQPort:                os.Getenv("RABBITMQ_PORT"),
		RabbitMQUser:                os.Getenv("RABBITMQ_USER"),
		RabbitMQPassword:            os.Getenv("RABBITMQ_PASSWORD"),
		RabbitMQConsumeQueueName:    os.Getenv("RABBITMQ_CONSUME_QUEUE_NAME"),
		RabbitMQConsumeDLQQueueName: os.Getenv("RABBITMQ_CONSUME_DLQ_QUEUE_NAME"),
	}
}
