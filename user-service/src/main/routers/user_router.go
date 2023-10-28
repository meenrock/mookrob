package routers

import (
	services "github.com/mookrob/serviceuser/main/services/rest"
	authentication "github.com/mookrob/shared/authentication"
	constants "github.com/mookrob/shared/constants"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine, ctrls *services.UserRestService) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("/create", ctrls.CreateUser)
		userRoutes.GET("/detail", authentication.AuthMiddleware(constants.GENERAL_USER), ctrls.GetUserById)
		userRoutes.GET("/fav-food", authentication.AuthMiddleware(constants.GENERAL_USER), ctrls.GetUserFavFoodByUserId)
	}
}
