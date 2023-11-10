package repositories

import (
	"database/sql"
	"time"

	"github.com/mookrob/servicemeal/main/models"

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

func (r *UserMealRepository) CreateDailyUserMeal(dailyUserMeal models.DailyUserMeal) error {
	_, err := r.DB.Exec("INSERT INTO daily_user_meal ("+
		"meal_id, "+
		"user_id, "+
		"meal_time, "+
		"date, "+
		"created_at, "+
		"updated_at "+
		") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", dailyUserMeal.MealId, dailyUserMeal.UserId, dailyUserMeal.MealTime,
		dailyUserMeal.Date, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
