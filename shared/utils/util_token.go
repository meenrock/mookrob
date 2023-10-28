package utils

import (
	"log"

	"github.com/google/uuid"
	"github.com/mookrob/shared/constants"
	models "github.com/mookrob/shared/models"

	"github.com/dgrijalva/jwt-go"
)

func ExtractUserData(userDataRaw any) (models.UserData, bool) {
	mapClaims, ok := userDataRaw.(jwt.MapClaims)
	if ok != true {
		log.Printf("ExtractUserData: failed parse claims")
		return models.UserData{}, false
	}

	var role constants.Role
	ok = role.ParseString(mapClaims["role"].(string))
	if !ok {
		log.Printf("ExtractUserData: failed parse role")
		return models.UserData{}, false
	}
	userId, err := uuid.Parse(mapClaims["user_id"].(string))
	if err != nil {
		log.Printf("ExtractUserData: failed parse user id")
		return models.UserData{}, false
	}
	userName := mapClaims["user_name"].(string)

	userData := models.UserData{
		Role:     role,
		UserId:   userId,
		UserName: userName,
	}

	return userData, true
}
