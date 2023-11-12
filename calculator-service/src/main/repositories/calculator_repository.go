package repositories

import (
	"context"
	"database/sql"
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

func (r *UserCalculatorRepository) GetUserCalculationByUserId(id uuid.UUID) (*sql.Rows, error) {
	rows, err := r.DB.Query("SELECT "+
		"f.id, "+
		"f.name, "+
		"f.energy, "+
		"f.protein, "+
		"f.carbohydrate, "+
		"f.fat, "+
		"f.sodium, "+
		"f.cholesterol, "+
		"f.created_at, "+
		"f.updated_at "+
		"FROM user_food uf "+
		"inner join food f on uf.food_id = f.id "+
		"WHERE uf.user_id = $1 and uf.user_food_type = 'LIKE'", id)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *UserCalculatorRepositoryMongo) AddParameter(id uuid.UUID, collection string, db_name string, ctx *gin.Context) {
	coll := ConnectMongoDB().Database(db_name).Collection(collection)
	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal.", "detailed": err})
	}
	ctx.JSON(http.StatusOK, result)
}
