package repositories

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {

	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file (yaml, json, etc.)
	viper.AddConfigPath("../resources/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	MONGO_HOST := viper.GetString("mongo.host")
	MONGO_PORT := viper.GetString("mongo.port")
	MONGO_USER := viper.GetString("mongo.user")
	MONGO_PASSWORD := viper.GetString("mongo.password")

	mongoInfo := fmt.Sprintf("mongodb://%s:%s@%s:%s", MONGO_USER, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoInfo))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
	}
	return client
}
