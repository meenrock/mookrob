package queue

import (
	// client "restapi/proto/client"
	"log"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createRabbitMQConnection() (*amqp.Connection, *amqp.Channel) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env unable to load")
	}

	HOST := viper.GetString("rabbitmq.host")
	PORT := viper.GetString("rabbitmq.port")
	PWD := viper.GetString("rabbitmq.pwd")
	USER := viper.GetString("rabbitmq.user")

	dialStr := "amqp://" + USER + ":" + PWD + "@" + HOST + ":" + PORT + "/"

	conn, err := amqp.Dial(dialStr)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}
