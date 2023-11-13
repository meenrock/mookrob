package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	repositories "github.com/mookrob/servicecalculator/main/repositories"
	"github.com/mookrob/servicecalculator/main/routers"

	// mqtt_services "github.com/mookrob/servicecalculator/main/services/mq"
	rest_services "github.com/mookrob/servicecalculator/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
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

	go func() {
		repositories.ConnectMongoDB()
	}()

	go func() {
		authGrpcService := grpc_services.NewAuthenticationGrpcService(authRepository)
		pb.RegisterAuthServer(grpc_server, authGrpcService)

		grpc_port := fmt.Sprintf(":%v", GRPC_PORT)
		lis, err := net.Listen("tcp", grpc_port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		fmt.Println("GRPC Server listening on Port", grpc_port)
		if err := grpc_server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

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
