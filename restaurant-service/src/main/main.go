package main

import (
	"fmt"
	"log"

	routers "github.com/mookrob/servicerestaurant/main/routers"
	rest_services "github.com/mookrob/servicerestaurant/main/services/rest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load("../resources/.env")

	if err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}

	viper.AutomaticEnv()

	// DB connection
	REST_PORT := viper.GetString("REST_PORT")

	gin.SetMode(viper.GetString("GIN_MODE"))
	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())

	// create instances of services and controllers
	placeService := rest_services.NewPlaceRestService()
	routers.SetPlaceRoutes(gin_engine, placeService)

	// Start the server
	port := fmt.Sprintf(":%v", REST_PORT)
	fmt.Println("Server Running on Port", port)
	if err := gin_engine.Run(port); err != nil {
		log.Fatal(err)
	}
}
