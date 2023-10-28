package rest_services

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mookrob/servicemeal/main/models"
	"github.com/mookrob/servicemeal/main/repositories"
)

type UserFoodRestService struct {
	UserFoodRepository *repositories.UserFoodRepository
}

func NewUserFoodRestService(r *repositories.UserFoodRepository) *UserFoodRestService {
	return &UserFoodRestService{UserFoodRepository: r}
}

type UserFavFoodResponse struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Energy       float64  `json:"energy"`
	Protein      *float64 `json:"protein"`
	Carbohydrate *float64 `json:"carbohydrate"`
	Fat          *float64 `json:"fat"`
	Sodium       *float64 `json:"sodium"`
	Cholesterol  *float64 `json:"cholesterol"`
}

func (s *UserFoodRestService) GetUserFavFoodByUserId(ctx *gin.Context) {
	userData, exist := ctx.Get("userData")
	if exist != true {
		log.Printf("rest GetUserFavFoodByUserId failed parse userData")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	mapClaims, ok := userData.(jwt.MapClaims)
	if ok != true {
		log.Printf("rest GetUserFavFoodByUserId failed parse user id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	id, err := uuid.Parse(mapClaims["user_id"].(string))
	if err != nil {
		log.Printf("rest GetUserFavFoodByUserId failed parse user id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	rows, err := s.UserFoodRepository.GetUserFavFoodByUserId(id)
	if err != nil {
		log.Printf("rest GetUserFavFoodByUserId failed on user food repository call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer rows.Close()

	var foods []models.Food

	// Iterate through the result set and scan each row into a User model
	for rows.Next() {
		var food models.Food
		if err := rows.Scan(&food.Id, &food.Name, &food.Energy, &food.Protein, &food.Carbohydrate, &food.Fat,
			&food.Sodium, &food.Cholesterol, &food.CreatedAt, &food.UpdatedAt); err != nil {
			log.Printf("rest GetUserFavFoodByUserId failed on row scan: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		foods = append(foods, food)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("rest GetUserFavFoodByUserId failed query rows: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	ctx.JSON(http.StatusOK, foods)
	return
}
