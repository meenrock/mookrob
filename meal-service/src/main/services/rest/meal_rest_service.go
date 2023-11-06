package rest_services

import (
	"log"
	"net/http"

	"github.com/mookrob/servicemeal/main/models"
	"github.com/mookrob/servicemeal/main/repositories"

	"github.com/gin-gonic/gin"
)

type MealRestService struct {
	MealRepository *repositories.MealRepository
}

func NewMealRestService(r *repositories.MealRepository) *MealRestService {
	return &MealRestService{MealRepository: r}
}

// CreateMeal request model
type CreateMealRequest struct {
	Name         string   `json:"name" binding:"required"`
	Energy       float64  `json:"energy" binding:"required"`
	Protein      *float64 `json:"protein"`
	Carbohydrate *float64 `json:"carbohydrate"`
	Fat          *float64 `json:"fat"`
	Sodium       *float64 `json:"sodium"`
	Cholesterol  *float64 `json:"cholesterol"`
}

func (s *MealRestService) CreateMeal(ctx *gin.Context) {
	var request CreateMealRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("rest CreateMeal: error on parse request: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create meal
	newMeal := models.Meal{
		Name:         request.Name,
		Energy:       request.Energy,
		Protein:      request.Protein,
		Carbohydrate: request.Carbohydrate,
		Fat:          request.Fat,
		Sodium:       request.Sodium,
		Cholesterol:  request.Cholesterol,
	}
	err := s.MealRepository.CreateMeal(newMeal)
	if err != nil {
		log.Println("rest GetUserById: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}
