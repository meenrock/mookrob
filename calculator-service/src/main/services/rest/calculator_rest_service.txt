package services

import (
	"github.com/gin-gonic/gin"
	"github.com/mookrob/servicecalculator/main/repositories"
	mqtt_services "github.com/mookrob/servicecalculator/main/services/mq"
)

type UserCalculatorService struct {
	UserCalculatorRepository *repositories.UserCalculatorRepositoryMongo
}

func NewUserCalculatorRestService(r *repositories.UserCalculatorRepositoryMongo) *UserCalculatorService {
	return &UserCalculatorService{UserCalculatorRepository: r}
}

func (s *UserCalculatorService) GetUserCalculationByUserId(ctx *gin.Context) {
	conn, ch := mqtt_services.CreateRabbitMQConnection()
	defer conn.Close()
	defer ch.Close()
}

func (s *UserCalculatorService) CalculateUserParameter(ctx *gin.Context) {

}
