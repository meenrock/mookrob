package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/mookrob/serviceauth/main/enums"
)

type UserBodyData struct {
	ID            uuid.UUID `bson:"_id,omitempty"`
	UserId        uuid.UUID `bson:"user_id,omitempty"`
	TransactionId string    `bson:"transaction_id"`
	BMI           float64   `bson:"bmi"`
	BMR           float64   `bson:"bmi"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}

type User struct {
	Id          uuid.UUID
	Status      enums.Status
	FirstName   string
	LastName    string
	NickName    string
	PhoneNumber *string
	Email       string
	Gender      string
	Age         int64
	Height      float64
	Weight      float64
	ExpectedBmi *float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
