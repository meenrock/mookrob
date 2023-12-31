package routers

import (
	services "github.com/mookrob/servicemeal/main/services/rest"
	authentication "github.com/mookrob/shared/authentication"
	constants "github.com/mookrob/shared/constants"

	"github.com/gin-gonic/gin"
)

func SetMealRoutes(router *gin.Engine, ctrls *services.MealRestService) {
	mealRoutes := router.Group("/api/meal")
	{
		mealRoutes.POST("/create", authentication.AuthMiddleware(constants.ADMIN), ctrls.CreateMeal)
		mealRoutes.GET("/", authentication.AuthMiddleware(constants.ADMIN), ctrls.GetMealList)
		mealRoutes.PUT("/edit", authentication.AuthMiddleware(constants.ADMIN), ctrls.EditMeal)
		mealRoutes.GET("/suggest", authentication.AuthMiddleware(constants.GENERAL_USER), ctrls.SuggestMeal)
	}
}
