package routers

import (
	services "github.com/mookrob/servicecalculator/main/services/rest"

	"github.com/gin-gonic/gin"
)

func SetUserCalculatorRoutes(router *gin.Engine, ctrls *services.UserCalculatorService) {
	calRoutes := router.Group("/api/calculator")
	{
		calRoutes.GET("/test", ctrls.GetUserCalculationByUserId)
	}
}
