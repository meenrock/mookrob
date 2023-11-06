package routers

import (
	services "github.com/mookrob/serviceauth/main/services/rest"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.Engine, ctrls *services.AuthenticationRestService) {
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login", ctrls.Login)
	}
}
