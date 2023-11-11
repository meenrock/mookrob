package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	repositories "github.com/mookrob/servicecalculator/main/repositories"
	"github.com/mookrob/servicecalculator/main/routers"

	// mqtt_services "github.com/mookrob/servicecalculator/main/services/mq"
	rest_services "github.com/mookrob/servicecalculator/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())

	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file (yaml, json, etc.)
	viper.AddConfigPath("../resources/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	// DB connection
	DB_HOST := viper.GetString("database.host")
	DB_PORT := viper.GetString("database.port")
	DB_NAME := viper.GetString("database.name")
	DB_USER := viper.GetString("database.user")
	DB_PASSWORD := viper.GetString("database.password")
	PORT := viper.GetString("server.rest-port")
	MONGO_HOST := viper.GetString("mongo.host")
	MONGO_PORT := viper.GetString("mongo.port")
	MONGO_USER := viper.GetString("mongo.user")
	MONGO_PASSWORD := viper.GetString("mongo.password")

	// connect postgres
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// create instances of services and controllers
	calculatorRepository := repositories.NewUserCalculatorRepository(db)
	// userService := services.NewUserService(userRepository)
	// routers.SetUserRoutes(r, userService)
	mongoInfo := fmt.Sprintf("mongodb://%s:%s@%s:%s", MONGO_USER, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoInfo))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return
	}

	go func() {
		// Start the server

		userCalculatorService := rest_services.NewUserCalculatorRestService(calculatorRepository)
		routers.SetUserCalculatorRoutes(gin_engine, userCalculatorService)

		rest_port := fmt.Sprintf(":%v", PORT)
		fmt.Println("Server Running on Port", rest_port)
		if err := gin_engine.Run(rest_port); err != nil {
			log.Fatal(err)
		}
	}()

	select {}

}
