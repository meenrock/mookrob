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
		body, _ := c.GetRawData()
		log.Printf("Failed to parse request body: %v. Request body: %s", err, body)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := r.UserService.SaveUserBodyData(&userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user data", "detail": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User data saved successfully"})
}
