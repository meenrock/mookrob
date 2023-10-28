package database

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateRabbitMQConnection() (*amqp.Connection, *amqp.Channel) {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file (yaml, json, etc.)
	viper.AddConfigPath("../resources/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	HOST := viper.GetString("rabbitmq.host")
	PORT := viper.GetString("rabbitmq.port")
	USER := viper.GetString("rabbitmq.user")
	PASSWORD := viper.GetString("rabbitmq.password")

	dialStr := "amqp://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/"

	conn, err := amqp.Dial(dialStr)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}
