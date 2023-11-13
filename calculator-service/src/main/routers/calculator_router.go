package routers

import (
	services "github.com/mookrob/servicecalculator/main/services/rest"

	rabbitmq "github.com/mookrob/servicecalculator/main/rabbitmq"

	"github.com/gin-gonic/gin"
)

func SetUserCalculatorRoutes(router *gin.Engine, ctrls *services.UserCalculatorService, rabbitmq *rabbitmq.RabbitMQ) {
	calRoutes := router.Group("/api/calculator")
	{
		calRoutes.GET("/test", ctrls.GetUserCalculationByUserId)
		calRoutes.POST("/createparameter")
	}
}
