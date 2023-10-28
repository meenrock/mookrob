package models

import (
	constants "github.com/mookrob/shared/constants"

	"github.com/google/uuid"
)

type UserData struct {
	Role     constants.Role
	UserId   uuid.UUID
	UserName string
}
