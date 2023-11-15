package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/mookrob/serviceauth/main/grpc-server"
	repositories "github.com/mookrob/serviceauth/main/repositories"
	routers "github.com/mookrob/serviceauth/main/routers"
	grpc_services "github.com/mookrob/serviceauth/main/services/grpc"
	rest_services "github.com/mookrob/serviceauth/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	GRPC_PORT := viper.GetString("GRPC_PORT")

	gin.SetMode(viper.GetString("GIN_MODE"))
	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())

	grpc_server := grpc.NewServer()

	// // connect postgres
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// create instances of services and controllers
	authRepository := repositories.NewAuthenticationRepository(db)

	// run async?? server
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

	// run async?? server
	go func() {
		authRestService := rest_services.NewAuthenticationRestService(authRepository)
		routers.SetAuthRoutes(gin_engine, authRestService)

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
