package routers

import (
	services "github.com/mookrob/servicerestaurant/main/services/rest"
	authentication "github.com/mookrob/shared/authentication"
	constants "github.com/mookrob/shared/constants"

	"github.com/gin-gonic/gin"
)

func SetPlaceRoutes(router *gin.Engine, ctrls *services.PlaceRestService) {
	placeRoutes := router.Group("/api/place")
	{
		placeRoutes.GET("/search", authentication.AuthMiddleware(constants.GENERAL_USER), ctrls.SearchPlaces)
	}
}
