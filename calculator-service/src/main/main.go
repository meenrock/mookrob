package main

import (
	"fmt"
	"log"

	mqtt_services "github.com/mookrob/servicecalculator/main/services/mq"

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

	// DB connection
	// DB_HOST := viper.GetString("database.host")
	// DB_PORT := viper.GetString("database.port")
	// DB_NAME := viper.GetString("database.name")
	// DB_USER := viper.GetString("database.user")
	// DB_PASSWORD := viper.GetString("database.password")
	PORT := viper.GetString("server.port")

	// // connect postgres
	// psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// db, err := sql.Open("pgx", psqlInfo)
	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }

	// create instances of services and controllers
	// userRepository := repositories.NewUserRepository(db)
	// userService := services.NewUserService(userRepository)
	// routers.SetUserRoutes(r, userService)

	// Start the server
	port := fmt.Sprintf(":%v", PORT)
	fmt.Println("Server Running on Port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}

	go func() {
		conn, ch := mqtt_services.CreateRabbitMQConnection()
		defer conn.Close()
		defer ch.Close()

	}()

}
