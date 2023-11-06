package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	repositories "github.com/mookrob/servicemeal/main/repositories"
	routers "github.com/mookrob/servicemeal/main/routers"
	grpc_services "github.com/mookrob/servicemeal/main/services/grpc"
	rest_services "github.com/mookrob/servicemeal/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	pb "github.com/mookrob/servicemeal/main/grpc-server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())

	grpc_server := grpc.NewServer()

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
	REST_PORT := viper.GetString("server.rest-port")
	GRPC_PORT := viper.GetString("server.grpc-port")

	// connect postgres
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// create instances of services and controllers
	userMealRepository := repositories.NewUserMealRepository(db)

	// run async?? server
	go func() {
		userMealGrpcService := grpc_services.NewUserMealGrpcService(userMealRepository)
		pb.RegisterMealServer(grpc_server, userMealGrpcService)

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

	// run async?? server
	go func() {
		userMealRestService := rest_services.NewUserMealRestService(userMealRepository)
		routers.SetUserMealRoutes(gin_engine, userMealRestService)

		// Start the server
		rest_port := fmt.Sprintf(":%v", REST_PORT)
		fmt.Println("Rest Server Running on Port", rest_port)
		if err := gin_engine.Run(rest_port); err != nil {
			log.Fatal(err)
		}
	}()

	// keep server running
	select {}

}
