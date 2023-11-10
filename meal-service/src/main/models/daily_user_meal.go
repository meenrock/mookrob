package models

import (
	"time"

	"github.com/mookrob/servicemeal/main/enums"

	"github.com/google/uuid"
)

type DailyUserMeal struct {
	Id        uuid.UUID      `gorm:"column:id; primary_key default:uuid_generate_v4()"`
	MealId    uuid.UUID      `gorm:"column:meal_id"`
	UserId    uuid.UUID      `gorm:"column:user_id"`
	MealTime  enums.MealTime `gorm:"column:meal_time; uniqueIndex:mealtime_date_index"`
	Date      time.Time      `gorm:"column:date; uniqueIndex:mealtime_date_index"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (t *DailyUserMeal) TableName() string {
	return "daily_user_meal"
}
