package repositories

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	// "go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCalculatorRepository struct {
	DB *sql.DB
}

type UserCalculatorRepositoryMongo struct {
	MONGO *mongo.Database
}

func NewUserCalculatorRepository(db *sql.DB) *UserCalculatorRepository {
	return &UserCalculatorRepository{DB: db}
}

func (r *UserCalculatorRepository) GetUserCalculationByUserId(id uuid.UUID) (mongo.InsertOneResult, error) {
	coll := ConnectMongoDB().Database("db_name").Collection("collection")
	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{})
	if err != nil {
		log.Fatal("Error finding document:", err)
		return *result, err
	}

	return *result, nil
}

func (r *UserCalculatorRepository) GetUserCalculationBMI(id uuid.UUID) (mongo.SingleResult, error) {
	coll := ConnectMongoDB().Database("db_name").Collection("collection")

	filter := bson.D{{}}

	result := coll.FindOne(context.Background(), filter)
	if result.Err() != nil {
		log.Fatal("Error finding document:", result.Err())
		return mongo.SingleResult{}, result.Err()
	}

	return *result, nil
}

func (r *UserCalculatorRepositoryMongo) AddParameter(id uuid.UUID, collection string, db_name string, ctx *gin.Context) {
	coll := ConnectMongoDB().Database(db_name).Collection(collection)
	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal.", "detailed": err})
	}
	ctx.JSON(http.StatusOK, result)
}
