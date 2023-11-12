package services

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mookrob/servicecalculator/main/repositories"

	// pub "github.com/mookrob/serviceuser/main/queue"
	// repositories "github.com/mookrob/serviceuser/main/repositories"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type UserCalculatorService struct {
	UserCalculatorRepository *repositories.UserCalculatorRepository
}

func PublishUserInfo(ctx *gin.Context) {
	// conn, ch := CreateRabbitMQConnection()

	// req := &pb.GetUserFavFoodRequest{
	// 	Id: ctx.Param("id"),
	// }

	// pub.PublishMessage(ch, "", message)

}

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

func (s *UserCalculatorService) ProduceUserInformationMessage(ctx *gin.Context) {

}

func ConsumeMessage(ch *amqp.Channel, QueueName string) {

	q, err := ch.QueueDeclare(
		QueueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Printf(" [*] Waiting for messages in %s. To exit press CTRL+C\n", QueueName)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received")
			log.Printf("%s", d.Body)

			// Simulate processing time
			secs := bytes.Count(d.Body, []byte("."))
			time.Sleep(time.Duration(secs) * time.Second)

			log.Printf(" [x] Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
