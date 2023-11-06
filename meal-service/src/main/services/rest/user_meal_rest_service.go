package rest_services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mookrob/servicemeal/main/models"
	"github.com/mookrob/servicemeal/main/repositories"
	"github.com/mookrob/shared/utils"
)

type UserMealRestService struct {
	UserMealRepository *repositories.UserMealRepository
}

func NewUserMealRestService(r *repositories.UserMealRepository) *UserMealRestService {
	return &UserMealRestService{UserMealRepository: r}
}

type UserFavMealResponse struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Energy       float64  `json:"energy"`
	Protein      *float64 `json:"protein"`
	Carbohydrate *float64 `json:"carbohydrate"`
	Fat          *float64 `json:"fat"`
	Sodium       *float64 `json:"sodium"`
	Cholesterol  *float64 `json:"cholesterol"`
}

func (s *UserMealRestService) GetUserFavMealByUserId(ctx *gin.Context) {
	userDataRaw, exist := ctx.Get("userData")
	if exist != true {
		log.Printf("rest GetUserFavMealByUserId failed parse userData")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	userData, ok := utils.ExtractUserData(userDataRaw)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	rows, err := s.UserMealRepository.GetUserFavMealByUserId(userData.UserId)
	if err != nil {
		log.Printf("rest GetUserFavMealByUserId failed on user meal repository call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer rows.Close()

	var meals []models.Meal

	// Iterate through the result set and scan each row into a User model
	for rows.Next() {
		var meal models.Meal
		if err := rows.Scan(&meal.Id, &meal.Name, &meal.Energy, &meal.Protein, &meal.Carbohydrate, &meal.Fat,
			&meal.Sodium, &meal.Cholesterol, &meal.CreatedAt, &meal.UpdatedAt); err != nil {
			log.Printf("rest GetUserFavMealByUserId failed on row scan: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		meals = append(meals, meal)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("rest GetUserFavMealByUserId failed query rows: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	ctx.JSON(http.StatusOK, meals)
	return
}
