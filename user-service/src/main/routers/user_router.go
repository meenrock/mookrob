package routers

import (
	services "github.com/mookrob/serviceuser/main/services/rest"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine, ctrls *services.UserRestService) {
	userRoutes := router.Group("/api/users")
	{
		userRoutes.GET("/:id", ctrls.GetUserById)
		userRoutes.GET("/fav-food/:id", ctrls.GetUserFavFoodByUserId)
	}
}
