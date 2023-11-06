package repositories

import (
	"database/sql"

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