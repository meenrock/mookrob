package config

import (
	"time"

	"github.com/mookrob/servicecalculator/main/queue"
	"github.com/mookrob/servicecalculator/main/rabbitmq"
)

type Config struct {
	HTTPAddress string
	RabbitMQ    rabbitmq.Config
	UserAMQP    queue.AMQPConfig
}

func New() Config {
	var cnf Config

	cnf.HTTPAddress = ":8080"

	cnf.RabbitMQ.Schema = "amqp"
	cnf.RabbitMQ.Username = "meendev"
	cnf.RabbitMQ.Password = ""
	cnf.RabbitMQ.Host = "192.168.1.187"
	cnf.RabbitMQ.Port = "5672"
	cnf.RabbitMQ.Vhost = "my_app"
	cnf.RabbitMQ.ConnectionName = "MY_APP"
	cnf.RabbitMQ.ChannelNotifyTimeout = 100 * time.Millisecond
	cnf.RabbitMQ.Reconnect.Interval = 500 * time.Millisecond
	cnf.RabbitMQ.Reconnect.MaxAttempt = 7200

	cnf.UserAMQP.Create.ExchangeName = "user"
	cnf.UserAMQP.Create.ExchangeType = "direct"
	cnf.UserAMQP.Create.RoutingKey = "create"
	cnf.UserAMQP.Create.QueueName = "parameter_create"

	return cnf
}
