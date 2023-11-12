package main

import (
	"fmt"
	"log"

	routers "github.com/mookrob/servicerestaurant/main/routers"
	rest_services "github.com/mookrob/servicerestaurant/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file (yaml, json, etc.)
	viper.AddConfigPath("../resources/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	PORT := viper.GetString("server.port")

	// create instances of services and controllers
	placeService := rest_services.NewPlaceRestService()
	routers.SetPlaceRoutes(r, placeService)

	// Start the server
	port := fmt.Sprintf(":%v", PORT)
	fmt.Println("Server Running on Port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
