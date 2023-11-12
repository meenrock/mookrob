package main

import (
	"database/sql"
	"fmt"
	"log"

	repositories "github.com/mookrob/serviceuser/main/repositories"
	routers "github.com/mookrob/serviceuser/main/routers"
	rest_services "github.com/mookrob/serviceuser/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
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
	DB_HOST := viper.GetString("DB_HOST")
	DB_PORT := viper.GetString("DB_PORT")
	DB_NAME := viper.GetString("DB_NAME")
	DB_USER := viper.GetString("DB_USERNAME")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")
	REST_PORT := viper.GetString("REST_PORT")

	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())
	gin.SetMode(viper.GetString("GIN_MODE"))

	// connect postgres
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// create instances of services and controllers
	userRepository := repositories.NewUserRepository(db)

	go func() {
		userRestService := rest_services.NewUserRestService(userRepository)
		routers.SetUserRoutes(gin_engine, userRestService)

		// Start the server
		port := fmt.Sprintf(":%v", REST_PORT)
		fmt.Println("Server Running on Port", REST_PORT)
		if err := gin_engine.Run(port); err != nil {
			log.Fatal(err)
		}
	}()

	select {}

}
