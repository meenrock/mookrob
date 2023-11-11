package services

import (
	"fmt"
	"log"

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
