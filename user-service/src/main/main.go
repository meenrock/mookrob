package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	repositories "github.com/mookrob/serviceuser/main/repositories"
	routers "github.com/mookrob/serviceuser/main/routers"
	"google.golang.org/grpc"

	grpc_services "github.com/mookrob/serviceuser/main/services/grpc"
	rest_services "github.com/mookrob/serviceuser/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"

	pb "github.com/mookrob/serviceuser/main/grpc-server/github.com/mookrob/serviceuser"
	"github.com/spf13/viper"
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
	userRepository := repositories.NewUserRepository(db)

	go func() {
		userGrpcService := grpc_services.NewUserGrpcService(userRepository)
		pb.RegisterUserServer(grpc_server, userGrpcService)

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

	userRestService := rest_services.NewUserRestService(userRepository)
	routers.SetUserRoutes(gin_engine, userRestService)

	// Start the server
	port := fmt.Sprintf(":%v", REST_PORT)
	fmt.Println("Server Running on Port", REST_PORT)
	if err := gin_engine.Run(port); err != nil {
		log.Fatal(err)
	}
}
