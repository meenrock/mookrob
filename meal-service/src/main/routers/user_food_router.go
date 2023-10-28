package routers

import (
	services "github.com/mookrob/servicemeal/main/services/rest"
	authentication "github.com/mookrob/shared/authentication"
	constants "github.com/mookrob/shared/constants"

	"github.com/gin-gonic/gin"
)

func SetUserFoodRoutes(router *gin.Engine, ctrls *services.UserFoodRestService) {
	userRoutes := router.Group("/api/meal/user-food")
	{
		userRoutes.GET("/fav", authentication.AuthMiddleware(constants.USER), ctrls.GetUserFavFoodByUserId)
	}
}
