package queue

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func PublishMessage(ch *amqp.Channel, QueueName string, message string) {

	ctx := context.Background()
	err := ch.PublishWithContext(
		ctx,
		"",        // exchange
		QueueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	failOnError(err, "Failed to publish a message")

	fmt.Printf(" [x] Sent: %s\n", message)
}
