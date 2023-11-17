package models

import (
	"time"

	"github.com/google/uuid"
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
