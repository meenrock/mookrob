package main

import (
	"context"
	"fmt"
	"log"
	"net"

	repositories "github.com/mookrob/servicecalculator/main/repositories"
	router "github.com/mookrob/servicecalculator/main/routers"

	mqtt_services "github.com/mookrob/servicecalculator/main/services/mq"

	// rest_services "github.com/mookrob/servicecalculator/main/services/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
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
	// PORT := viper.GetString("server.rest-port")
	GRPC_PORT := viper.GetString("server.grpc-port")
	MONGO_HOST := viper.GetString("mongo.host")
	MONGO_PORT := viper.GetString("mongo.port")
	MONGO_USER := viper.GetString("mongo.user")
	MONGO_PASSWORD := viper.GetString("mongo.password")
	MONGO_DATABASE := viper.GetString("mongo.database")
	MongoDBConnectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", MONGO_USER, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT)

	// connect postgres
	// psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// db, err := sql.Open("pgx", psqlInfo)
	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }

	// db := repositories.ConnectMongoDB()

	// create instances of services and controllers
	// calculatorRepository := repositories.NewUserCalculatorRepository()
	// userService := services.NewUserService(userRepository)
	// routers.SetUserRoutes(r, userService)

	go func() {

		client, err := mongo.NewClient(options.Client().ApplyURI(MongoDBConnectionString))
		if err != nil {
			log.Fatalf("Failed to create MongoDB client: %v", err)
		}
		if err := client.Connect(context.Background()); err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		database := client.Database(MONGO_DATABASE)

		CalculatorRepository := repositories.NewUserCalculatorRepository(database)

		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				log.Fatalf("Failed to disconnect MongoDB client: %v", err)
			}
		}()

		lis, err := net.Listen("tcp", GRPC_PORT)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpc_server := grpc.NewServer()
		CalculatorRepository := repositories.NewUserCalculatorRepository()

	}()

	go func() {
		// Start the server

		// userCalculatorService := rest_services.NewUserCalculatorRestService(calculatorRepository)
		// routers.SetUserCalculatorRoutes(gin_engine, userCalculatorService)

		// rest_port := fmt.Sprintf(":%v", PORT)
		// fmt.Println("Server Running on Port", rest_port)
		// if err := gin_engine.Run(rest_port); err != nil {
		// 	log.Fatal(err)
		// }
	}()

	go func() {
		mongoDBRepo, err := repositories.NewMongoDBRepository("your_mongo_connection_string", "your_db_name", "your_collection_name")
		if err != nil {
			log.Fatalf("Failed to initialize MongoDB repository: %v", err)
		}

		sqsRepo, err := repositories.NewSQSRepository("your_aws_region", "your_sqs_queue_url")
		if err != nil {
			log.Fatalf("Failed to initialize SQS repository: %v", err)
		}

		userService := mqtt_services.NewUserService(mongoDBRepo, sqsRepo)
		appRouter := router.NewRouter(userService)
		engine := appRouter.SetupRoutes()

		engine.Run(":8080")
	}()

	select {}

}
