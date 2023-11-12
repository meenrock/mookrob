package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

// album represents data about a record album.
type meal struct {
	ID   string `json:"id"`
	Menu string `json:"title"`
}

// albums slice to seed record album data.
var meals = []meal{
	{ID: "1", Menu: "Congee"},
	{ID: "2", Menu: "Padthai"},
	{ID: "3", Menu: "Noodle"},
}

// Google Maps API key.
const googleMapsAPIKey = "AIzaSyD1T-WZxo0qHDTEexvV-CWJg0nLFRz5nDM"

func main() {
	router := gin.Default()
	router.GET("/meals", getMeal)
	router.GET("/meals/:id", getMealByID)
	router.POST("/meals", postMeal)
	router.GET("/search", searchPlaces)

	router.Run("localhost:8080")
}

// getMeals responds with the list of all albums as JSON.
func getMeal(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, meals)
}

// postAlbums adds an album from JSON received in the request body.
func postMeal(c *gin.Context) {
	var newMeal meal

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newMeal); err != nil {
		return
	}

	// Add the new album to the slice.
	meals = append(meals, newMeal)
	c.IndentedJSON(http.StatusCreated, newMeal)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getMealByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range meals {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "meal not found"})
}

// searchPlaces searches for places on Google Maps using the input from GET.
func searchPlaces(c *gin.Context) {
	input := c.Query("input")

	// Create a new Google Maps client.
	client, err := maps.NewClient(maps.WithAPIKey(googleMapsAPIKey))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Create a new text search request.
	req := &maps.TextSearchRequest{
		Query: input, // Set the Query field with the input from the request.
	}

	// Send the search request.
	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return the found places to the client.
	c.IndentedJSON(http.StatusOK, resp.Results)
}
