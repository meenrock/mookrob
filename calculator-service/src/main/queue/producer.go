package queue

import (
	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQProducer() (*RabbitMQProducer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	err.failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	err.failOnError(err, "Failed to open a channel")
	defer ch.Close()

	return &RabbitMQProducer{
		conn:    conn,
		channel: ch,
	}, nil
}

func (c *RabbitMQProducer) PublishMessage() {

}
