package rest_services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mookrob/servicemeal/main/enums"
	"github.com/mookrob/servicemeal/main/models"
	"github.com/mookrob/servicemeal/main/repositories"
	"github.com/mookrob/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if !exist {
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
}

type CreateDailyUserMealRequest struct {
	MealId   uuid.UUID      `json:"meal_id" binding:"required"`
	MealTime enums.MealTime `json:"meal_time" binding:"required"`
}

func (s *UserMealRestService) CreateDailyUserMeal(ctx *gin.Context) {
	var request CreateDailyUserMealRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("rest CreateDailyUserMeal: error on parse request: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDataRaw, exist := ctx.Get("userData")
	if !exist {
		log.Println("rest CreateDailyUserMeal: failed parse userData")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	userData, ok := utils.ExtractUserData(userDataRaw)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	existId, _ := s.UserMealRepository.FindIdByUserIdMealTimeDate(userData.UserId, request.MealTime, time.Now())

	// create daily user meal
	newDailyUserMeal := models.DailyUserMeal{
		MealId:   request.MealId,
		UserId:   userData.UserId,
		MealTime: request.MealTime,
		Date:     time.Now(),
	}

	if existId != nil {
		fmt.Println("rest CreateDailyUserMeal: Record update!")
		err := s.UserMealRepository.UpdateDailyUserMeal(*existId, newDailyUserMeal)
		if err != nil {
			log.Println("rest CreateDailyUserMeal: failed on user meal UpdateDailyUserMeal call: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}
	} else {
		fmt.Println("rest CreateDailyUserMeal: Record create!")
		err := s.UserMealRepository.CreateDailyUserMeal(newDailyUserMeal)
		if err != nil {
			log.Println("rest CreateDailyUserMeal: failed on user meal CreateDailyUserMeal call: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}

type DailyUserMealResponse struct {
	MealId       uuid.UUID      `json:"meal_id"`
	Name         string         `json:"name"`
	Energy       float64        `json:"energy"`
	Protein      *float64       `json:"protein"`
	Carbohydrate *float64       `json:"carbohydrate"`
	Fat          *float64       `json:"fat"`
	Sodium       *float64       `json:"sodium"`
	Cholesterol  *float64       `json:"cholesterol"`
	MealTime     enums.MealTime `json:"meal_time"`
	Date         time.Time      `json:"date"`
}

func (s *UserMealRestService) GetDailyUserMeal(ctx *gin.Context) {
	userDataRaw, exist := ctx.Get("userData")
	if !exist {
		log.Println("rest GetDailyUserMeal: failed parse userData")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	userData, ok := utils.ExtractUserData(userDataRaw)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	rows, err := s.UserMealRepository.GetDailyUserMealByUserIdAndDate(userData.UserId, time.Now())
	if err != nil {
		log.Println("rest GetDailyUserMeal failed on user meal repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer rows.Close()

	var responseList []DailyUserMealResponse

	// Iterate through the result set and scan each row into a User model
	for rows.Next() {
		var meal models.Meal
		var mealTime enums.MealTime
		var date time.Time
		if err := rows.Scan(&meal.Id, &meal.Name, &meal.Energy, &meal.Protein, &meal.Carbohydrate, &meal.Fat,
			&meal.Sodium, &meal.Cholesterol, &meal.CreatedAt, &meal.UpdatedAt, &mealTime, &date); err != nil {
			log.Println("rest GetDailyUserMeal failed on row scan: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		dailyUserMealResponse := DailyUserMealResponse{
			MealId:       meal.Id,
			Name:         meal.Name,
			Energy:       meal.Energy,
			Protein:      meal.Protein,
			Carbohydrate: meal.Carbohydrate,
			Fat:          meal.Fat,
			Sodium:       meal.Sodium,
			Cholesterol:  meal.Cholesterol,
			MealTime:     mealTime,
			Date:         date,
		}

		responseList = append(responseList, dailyUserMealResponse)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Println("rest GetDailyUserMeal failed query rows: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	ctx.JSON(http.StatusOK, responseList)
}
