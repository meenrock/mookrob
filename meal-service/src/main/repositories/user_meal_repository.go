package repositories

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserMealRepository struct {
	DB *sql.DB
}

func NewUserMealRepository(db *sql.DB) *UserMealRepository {
	return &UserMealRepository{DB: db}
}

func (r *UserMealRepository) GetUserFavMealByUserId(id uuid.UUID) (*sql.Rows, error) {
	rows, err := r.DB.Query("SELECT "+
		"m.id, "+
		"m.name, "+
		"m.energy, "+
		"m.protein, "+
		"m.carbohydrate, "+
		"m.fat, "+
		"m.sodium, "+
		"m.cholesterol, "+
		"m.created_at, "+
		"m.updated_at "+
		"FROM user_meal um "+
		"inner join meal m on um.meal_id = m.id "+
		"WHERE um.user_id = $1 and um.user_meal_type = 'LIKE'", id)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
