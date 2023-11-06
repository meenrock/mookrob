package repositories

import (
	"database/sql"
	"time"

	"github.com/mookrob/servicemeal/main/models"
)

type MealRepository struct {
	DB *sql.DB
}

func NewMealRepository(db *sql.DB) *MealRepository {
	return &MealRepository{DB: db}
}

func (r *MealRepository) CreateMeal(meal models.Meal) error {
	_, err := r.DB.Exec("INSERT INTO meal ("+
		"name, "+
		"energy, "+
		"protein, "+
		"carbohydrate, "+
		"fat, "+
		"sodium, "+
		"cholesterol, "+
		"created_at, "+
		"updated_at "+
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", meal.Name, meal.Energy, meal.Protein, meal.Carbohydrate,
		meal.Fat, meal.Sodium, meal.Cholesterol, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
