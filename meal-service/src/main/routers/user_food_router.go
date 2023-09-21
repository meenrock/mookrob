package routers

import (
	services "github.com/mookrob/servicemeal/main/services/rest"

	"github.com/gin-gonic/gin"
)

func SetUserFoodRoutes(router *gin.Engine, ctrls *services.UserFoodRestService) {
	userRoutes := router.Group("/api/meal/user-food")
	{
		userRoutes.GET("/fav/:id", ctrls.GetUserFavFoodByUserId)
	}
}
