package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type RabbitMQProducer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQProducer() (*RabbitMQProducer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	return &RabbitMQProducer{
		conn:    conn,
		channel: ch,
	}, nil
}

func (c *RabbitMQProducer) PublishMessage(message []byte) error {
	if c.channel == nil {
		return fmt.Errorf("channel is not open")
	}

	queueName := "my-queue"

	// Declare the queue (you should declare it only once, ideally outside of this function)
	_, err := c.channel.QueueDeclare(
		queueName,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue: %v", err)
	}

	// Publish a message to the queue
	err = c.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to publish a message: %v", err)
	}

	log.Printf("Sent: %s", message)

	return nil
}
