package repositories

import (
	"database/sql"
	"time"

	"github.com/mookrob/servicemeal/main/enums"
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
		"inner join meals m on um.meal_id = m.id "+
		"WHERE um.user_id = $1 and um.user_meal_type = 'LIKE'", id)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *UserMealRepository) FindIdByUserIdMealTimeDate(userId uuid.UUID, mealTime enums.MealTime, date time.Time) (*uuid.UUID, error) {
	var id uuid.UUID
	err := r.DB.QueryRow("SELECT id FROM daily_user_meal WHERE user_id = $1 AND meal_time = $2 AND date = $3", userId, mealTime, date).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *UserMealRepository) UpdateDailyUserMeal(existId uuid.UUID, dailyUserMeal models.DailyUserMeal) error {
	_, err := r.DB.Exec("UPDATE daily_user_meal "+
		"SET meal_id = $1, "+
		"updated_at = $2 "+
		"WHERE id = $3", dailyUserMeal.MealId, time.Now(), existId)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserMealRepository) CreateDailyUserMeal(dailyUserMeal models.DailyUserMeal) error {
	_, err := r.DB.Exec("INSERT INTO daily_user_meal ("+
		"meal_id, "+
		"user_id, "+
		"meal_time, "+
		"date, "+
		"created_at, "+
		"updated_at "+
		") VALUES ($1, $2, $3, $4, $5, $6)", dailyUserMeal.MealId, dailyUserMeal.UserId, dailyUserMeal.MealTime,
		dailyUserMeal.Date, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (r *UserMealRepository) GetDailyUserMealByUserIdAndDate(id uuid.UUID, date time.Time) (*sql.Rows, error) {
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
		"m.updated_at, "+
		"dm.meal_time, "+
		"dm.date "+
		"FROM daily_user_meal dm "+
		"inner join meals m on dm.meal_id = m.id "+
		"WHERE dm.user_id = $1 and dm.date = $2", id, date.Format("2006-01-02"))

	if err != nil {
		return nil, err
	}

	return rows, nil
}
