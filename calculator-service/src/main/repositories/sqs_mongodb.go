package repositories

import (
	"context"
	"log"
	"time"

	"github.com/mookrob/servicecalculator/main/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDBRepository(connectionString, dbName, collectionName string) (*MongoDBRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &MongoDBRepository{
		Client:     client,
		Database:   db,
		Collection: collection,
	}, nil
}

func (repo *MongoDBRepository) InsertUserBodyData(userData *models.UserBodyData) error {

	userData.CreatedAt = time.Now()
	userData.UpdatedAt = time.Now()

	insertResult, err := repo.Collection.InsertOne(context.TODO(), userData)
	if err != nil {
		log.Printf("Failed to insert user data into MongoDB: %v", err)
		return err
	}

	log.Printf("Inserted user data with ID %v", insertResult.InsertedID)

	return nil
}
