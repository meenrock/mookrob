package routers

import (
	services "github.com/mookrob/servicemeal/main/services/rest"
	authentication "github.com/mookrob/shared/authentication"
	constants "github.com/mookrob/shared/constants"

	"github.com/gin-gonic/gin"
)

func SetUserMealRoutes(router *gin.Engine, ctrls *services.UserMealRestService) {
	userRoutes := router.Group("/api/meal/user-meal")
	{
		userRoutes.GET("/fav", authentication.AuthMiddleware(constants.GENERAL_USER), ctrls.GetUserFavMealByUserId)
	}
}
