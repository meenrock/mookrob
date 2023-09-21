package models

import (
	"time"

	"github.com/mookrob/serviceuser/main/enums"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID    `gorm:"column:id; primary_key default:uuid_generate_v4()"`
	Status      enums.Status `gorm:"column:status"`
	FirstName   string       `gorm:"column:first_name" orm:"size(100)"`
	LastName    string       `gorm:"column:last_name" orm:"size(100)"`
	NickName    string       `gorm:"column:nick_name" orm:"size(100)"`
	PhoneNumber *string      `gorm:"column:phone_number" orm:"size(10)"`
	Email       string       `gorm:"column:email; unique" orm:"size(100)"`
	Gender      string       `gorm:"column:gender" orm:"size(10)"`
	Age         int64        `gorm:"column:age"`
	Height      float64      `gorm:"column:height" orm:"size(150)"`
	Weight      float64      `gorm:"column:weight" orm:"size(150)"`
	ExpectedBmi *float64     `gorm:"column:expected_bmi"`
	CreatedAt   time.Time    `gorm:"column:created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"`
}

func (t *User) TableName() string {
	return "users"
}
