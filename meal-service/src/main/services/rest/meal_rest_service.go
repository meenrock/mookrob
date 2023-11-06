package rest_services

import (
	"log"
	"net/http"

	"github.com/google/uuid"
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

type MealListRequest struct {
	QueryString string `form:"query_string" `
	Page        *int   `form:"page" binding:"required"`
	PageSize    *int   `form:"page_size" binding:"required"`
}
type MealListResponse struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Energy       float64   `json:"energy"`
	Protein      *float64  `json:"protein"`
	Carbohydrate *float64  `json:"carbohydrate"`
	Fat          *float64  `json:"fat"`
	Sodium       *float64  `json:"sodium"`
	Cholesterol  *float64  `json:"cholesterol"`
}

func (s *MealRestService) GetMealList(ctx *gin.Context) {
	var request MealListRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Println("rest GetMealList: error on parse request: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call repo
	rows, err := s.MealRepository.GetMealList(request.QueryString, *request.Page, *request.PageSize)
	if err != nil {
		log.Println("rest GetMealDetail: failed on meal repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer rows.Close()

	var meals []MealListResponse

	// Iterate through the result set and scan each row into a User model
	for rows.Next() {
		var meal MealListResponse
		if err := rows.Scan(&meal.Id, &meal.Name, &meal.Energy, &meal.Protein, &meal.Carbohydrate, &meal.Fat,
			&meal.Sodium, &meal.Cholesterol); err != nil {
			log.Printf("rest GetMealDetail failed on row scan: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		meals = append(meals, meal)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("rest GetMealDetail failed query rows: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data_list": meals, "page": request.Page, "page_size": request.PageSize})
}
