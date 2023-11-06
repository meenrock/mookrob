package models

import (
	"time"

	"github.com/google/uuid"
)

type Meal struct {
	Id           uuid.UUID `gorm:"column:id; primary_key default:uuid_generate_v4()"`
	Name         string    `gorm:"column:name; unique"`
	Energy       float64   `gorm:"column:energy"`
	Protein      *float64  `gorm:"column:protein"`
	Carbohydrate *float64  `gorm:"column:carbohydrate"`
	Fat          *float64  `gorm:"column:fat"`
	Sodium       *float64  `gorm:"column:sodium"`
	Cholesterol  *float64  `gorm:"column:cholesterol"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (t *Meal) TableName() string {
	return "meal"
}
