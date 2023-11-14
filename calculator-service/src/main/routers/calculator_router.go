package routers

import (
	"net/http"

	mqtt "github.com/mookrob/servicecalculator/main/services/mq"
	services "github.com/mookrob/servicecalculator/main/services/rest"

	rabbitmq "github.com/mookrob/servicecalculator/main/rabbitmq"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{http.NewServeMux()}
}

func (rx *Router) SetUserCalculatorRoutes(router *gin.Engine, ctrls *services.UserCalculatorService, rabbitmq *rabbitmq.RabbitMQ) {
	createparameter := mqtt.NewCreate(rabbitmq)
	rx.HandleFunc("/parameter/add", createparameter.Handle)

	calRoutes := router.Group("/api/calculator")
	{
		calRoutes.GET("/test", ctrls.GetUserCalculationByUserId)
	}

}
