package models

import (
	"time"

	"github.com/mookrob/serviceauth/main/enums"
	constants "github.com/mookrob/shared/constants"

	"github.com/google/uuid"
)

type AuthenticationData struct {
	Id        uuid.UUID      `gorm:"column:id; primary_key default:uuid_generate_v4()"`
	Username  string         `gorm:"column:first_name unique" orm:"size(100)"`
	Password  string         `gorm:"column:first_name" orm:"size(255)"`
	UserId    *uuid.UUID     `gorm:"column:user_id; unique"`
	Role      constants.Role `gorm:"column:role"`
	Status    enums.Status   `gorm:"column:status"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (t *AuthenticationData) TableName() string {
	return "authentication_data"
}
