package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mookrob/servicecalculator/main/rabbitmq"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Create struct {
	rabbitmq *rabbitmq.RabbitMQ
}

func NewCreate(rabbitmq *rabbitmq.RabbitMQ) Create {
	return Create{
		rabbitmq: rabbitmq,
	}
}

func (c Create) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("ID")

	if err := c.publish(id); err != nil {
		log.Println(errors.Wrap(err, fmt.Sprintf("failed to create %s", id)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("created", id)
	w.WriteHeader(http.StatusAccepted)
}

func (c Create) publish(message string) error {
	channel, err := c.rabbitmq.Channel()
	if err != nil {
		return errors.Wrap(err, "failed to open channel")
	}
	defer channel.Close()

	if err := channel.Confirm(false); err != nil {
		return errors.Wrap(err, "failed to put channel in confirmation mode")
	}

	if err := channel.Publish(
		"parameter",
		"create",
		true,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			MessageId:    "A-UNIQUE-ID",
			ContentType:  "text/plain",
			Body:         []byte(message),
		},
	); err != nil {
		return errors.Wrap(err, "failed to publish message")
	}

	select {
	case ntf := <-channel.NotifyPublish(make(chan amqp.Confirmation, 1)):
		if !ntf.Ack {
			return errors.New("failed to deliver message to exchange/queue")
		}
	case <-channel.NotifyReturn(make(chan amqp.Return)):
		return errors.New("failed to deliver message to exchange/queue")
	case <-time.After(c.rabbitmq.ChannelNotifyTimeout):
		log.Println("message delivery confirmation to exchange/queue timed out")
	}

	return nil
}
