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
		log.Println("rest CreateMeal: failed on user repository call: ", err)
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

// EditMeal request model
type EditMealRequest struct {
	Id           uuid.UUID `json:"id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Energy       float64   `json:"energy" binding:"required"`
	Protein      *float64  `json:"protein"`
	Carbohydrate *float64  `json:"carbohydrate"`
	Fat          *float64  `json:"fat"`
	Sodium       *float64  `json:"sodium"`
	Cholesterol  *float64  `json:"cholesterol"`
}

func (s *MealRestService) EditMeal(ctx *gin.Context) {
	var request EditMealRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("rest EditMeal: error on parse request: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// edit meal
	editMeal := models.Meal{
		Id:           request.Id,
		Name:         request.Name,
		Energy:       request.Energy,
		Protein:      request.Protein,
		Carbohydrate: request.Carbohydrate,
		Fat:          request.Fat,
		Sodium:       request.Sodium,
		Cholesterol:  request.Cholesterol,
	}
	err := s.MealRepository.EditMeal(editMeal)
	if err != nil {
		log.Println("rest EditMeal: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}

// Meal Suggestion model
type MealSuggest struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type AllMealSuggest struct {
	Breakfast []MealSuggest `json:"breakfast"`
	Lunch     []MealSuggest `json:"lunch"`
	Dinner    []MealSuggest `json:"dinner"`
}

func (s *MealRestService) SuggestMeal(ctx *gin.Context) {

	caloriesPerDay := 1600

	// Calculate meal calories based on the desired percentages
	breakfastCalories := (caloriesPerDay * 3) / 10
	lunchCalories := (caloriesPerDay * 5) / 10
	dinnerCalories := (caloriesPerDay * 2) / 10

	log.Println("rest SuggestMeal: breakfastCalories: ", breakfastCalories)
	log.Println("rest SuggestMeal: lunchCalories: ", lunchCalories)
	log.Println("rest SuggestMeal: dinnerCalories: ", dinnerCalories)

	// Retrieve meal suggestions from the PostgreSQL database
	breakfastRows, err := s.MealRepository.GetSuggestMeal(breakfastCalories)
	if err != nil {
		log.Println("rest SuggestMeal: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	allSuggest := AllMealSuggest{
		Breakfast: make([]MealSuggest, 0),
		Lunch:     make([]MealSuggest, 0),
		Dinner:    make([]MealSuggest, 0),
	}

	for breakfastRows.Next() {
		var meal MealSuggest
		if err := breakfastRows.Scan(&meal.Id, &meal.Name); err != nil {
			log.Printf("rest SuggestMeal failed on row scan: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		allSuggest.Breakfast = append(allSuggest.Breakfast, meal)
	}

	lunchRows, err := s.MealRepository.GetSuggestMeal(lunchCalories)
	if err != nil {
		log.Println("rest SuggestMeal: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	if err != nil {
		log.Println("rest SuggestMeal: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	for lunchRows.Next() {
		var meal MealSuggest
		if err := lunchRows.Scan(&meal.Id, &meal.Name); err != nil {
			log.Printf("rest SuggestMeal failed on row scan: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		allSuggest.Lunch = append(allSuggest.Lunch, meal)
	}

	dinnerRows, err := s.MealRepository.GetSuggestMeal(dinnerCalories)
	if err != nil {
		log.Println("rest SuggestMeal: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	if err != nil {
		log.Println("rest SuggestMeal: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	for dinnerRows.Next() {
		var meal MealSuggest
		if err := dinnerRows.Scan(&meal.Id, &meal.Name); err != nil {
			log.Printf("rest SuggestMeal failed on row scan: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		allSuggest.Dinner = append(allSuggest.Dinner, meal)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": allSuggest})

}
