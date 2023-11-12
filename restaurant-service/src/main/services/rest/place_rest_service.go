package rest_services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"googlemaps.github.io/maps"
)

type PlaceRestService struct{}

func NewPlaceRestService() *PlaceRestService {
	return &PlaceRestService{}
}

// search place response
type PlaceResponse struct {
	FormattedAddress string           `json:"formatted_address"`
	Name             string           `json:"name"`
	Location         LocationResponse `json:"location"`
	PlaceId          string           `json:"place_id"`
	Rating           *float32         `json:"rating"`
	OpenNow          *bool            `json:"open_now"`
}
type LocationResponse struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (s *PlaceRestService) SearchPlaces(ctx *gin.Context) {
	mealName := ctx.Query("meal_name")

	// Create a new Google Maps client.
	client, err := maps.NewClient(maps.WithAPIKey(viper.GetString("google.api-key")))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Create a new text search request.
	req := &maps.TextSearchRequest{
		Query: mealName, // Set the Query field with the input from the request.
	}

	// Send the search request.
	resp, err := client.TextSearch(ctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var placeResponseList []PlaceResponse
	for _, value := range resp.Results {
		placeRes := PlaceResponse{
			FormattedAddress: value.FormattedAddress,
			Name:             value.Name,
			Location: LocationResponse{
				Lat: value.Geometry.Location.Lat,
				Lng: value.Geometry.Location.Lng,
			},
			PlaceId: value.PlaceID,
			Rating:  &value.Rating,
		}
		if value.OpeningHours != nil {
			placeRes.OpenNow = value.OpeningHours.OpenNow
		}

		placeResponseList = append(placeResponseList, placeRes)
	}

	// Return the found places to the client.
	ctx.IndentedJSON(http.StatusOK, placeResponseList)
}
