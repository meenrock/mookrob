package repositories

import (
	// "context"
	"database/sql"

	// "go.mongodb.org/mongo-driver/bson"
	// coll "go.mongodb.org/mongo-driver/mongo"

	"github.com/google/uuid"
)

type UserCalculatorRepository struct {
	DB *sql.DB
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

// func (r *UserCalculatorRepository) AddUserParameter(id uuid.UUID) {
// 	result, err := coll.InsertOne(
// 		context.TODO(),
// 		bson.D{
// 			{"item", "canvas"},
// 			{"qty", 100},
// 			{"tags", bson.A{"cotton"}},
// 			{"size", bson.D{
// 				{"h", 28},
// 				{"w", 35.5},
// 				{"uom", "cm"},
// 			}},
// 		})
// }
