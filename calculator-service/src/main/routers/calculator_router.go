package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mookrob/servicecalculator/main/models"
	services "github.com/mookrob/servicecalculator/main/services/mq"
)

type Router struct {
	UserService *services.UserService
}

func NewRouter(userService *services.UserService) *Router {
	return &Router{
		UserService: userService,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	calRoutes := router.Group("/api/calculator")
	{
		calRoutes.POST("/insert", r.handleSaveUserBodyData)
	}

	return router
}

func (r *Router) handleSaveUserBodyData(c *gin.Context) {
	var userData models.UserBodyData
	if err := c.BindJSON(&userData); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := r.UserService.SaveUserBodyData(&userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User data saved successfully"})
}

// import (
// 	"net/http"

// 	mqtt "github.com/mookrob/servicecalculator/main/services/mq"
// 	services "github.com/mookrob/servicecalculator/main/services/rest"

// 	rabbitmq "github.com/mookrob/servicecalculator/main/rabbitmq"

// 	"github.com/gin-gonic/gin"
// )

// type Router struct {
// 	*http.ServeMux
// }

// func NewRouter() *Router {
// 	return &Router{http.NewServeMux()}
// }

// func (rx *Router) SetUserCalculatorRoutes(router *gin.Engine, ctrls *services.UserCalculatorService, rabbitmq *rabbitmq.RabbitMQ) {
// 	createparameter := mqtt.NewCreate(rabbitmq)
// 	rx.HandleFunc("/parameter/add", createparameter.Handle)

// 	calRoutes := router.Group("/api/calculator")
// 	{
// 		calRoutes.GET("/test", ctrls.GetUserCalculationByUserId)
// 	}

// }
