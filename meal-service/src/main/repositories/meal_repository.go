package repositories

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mookrob/servicemeal/main/models"
)

type MealRepository struct {
	DB *sql.DB
}

func NewMealRepository(db *sql.DB) *MealRepository {
	return &MealRepository{DB: db}
}

func (r *MealRepository) CreateMeal(meal models.Meal) error {
	_, err := r.DB.Exec("INSERT INTO meal ("+
		"name, "+
		"energy, "+
		"protein, "+
		"carbohydrate, "+
		"fat, "+
		"sodium, "+
		"cholesterol, "+
		"created_at, "+
		"updated_at "+
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", meal.Name, meal.Energy, meal.Protein, meal.Carbohydrate,
		meal.Fat, meal.Sodium, meal.Cholesterol, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (r *MealRepository) GetMealById(id uuid.UUID) (models.Meal, error) {
	var meal models.Meal
	err := r.DB.QueryRow("SELECT "+
		"id, "+
		"name, "+
		"energy, "+
		"protein, "+
		"carbohydrate, "+
		"fat, "+
		"sodium, "+
		"cholesterol, "+
		"created_at, "+
		"updated_at "+
		"FROM meal "+
		"WHERE id = $1", id).Scan(&meal.Id, &meal.Name, &meal.Energy, &meal.Protein, &meal.Carbohydrate,
		&meal.Fat, &meal.Sodium, &meal.Cholesterol, &meal.CreatedAt, &meal.UpdatedAt)

	if err != nil {
		return models.Meal{}, err
	}

	return meal, nil
}

func (r *MealRepository) GetMealList(queryString string, page int, pageSize int) (*sql.Rows, error) {
	rows, err := r.DB.Query("SELECT "+
		"m.id, "+
		"m.name, "+
		"m.energy, "+
		"m.protein, "+
		"m.carbohydrate, "+
		"m.fat, "+
		"m.sodium, "+
		"m.cholesterol "+
		"FROM meals m "+
		"WHERE m.name LIKE '%' || $1 || '%' "+
		"ORDER BY m.name ASC "+
		"LIMIT $2 OFFSET $3", queryString, pageSize, (pageSize * page))

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *MealRepository) EditMeal(meal models.Meal) error {
	_, err := r.DB.Exec("UPDATE meals SET "+
		"name = $1, "+
		"energy = $2, "+
		"protein = $3, "+
		"carbohydrate = $4, "+
		"fat = $5, "+
		"sodium = $6, "+
		"cholesterol = $7, "+
		"updated_at = $8 "+
		"WHERE id = $9", meal.Name, meal.Energy, meal.Protein, meal.Carbohydrate,
		meal.Fat, meal.Sodium, meal.Cholesterol, time.Now(), meal.Id)

	if err != nil {
		return err
	}

	return nil
}

func (ctrls *MealRepository) SuggestMeal(c *gin.Context) {
	// Extract user's BMI, BMR, and desired calorie intake from the request
	bmi := c.PostForm("bmi")
	bmr := c.PostForm("bmr")
	desiredCalories := c.PostForm("desiredCalories")

	// Calculate calories per meal based on user's data
	caloriesPerMeal := calculateCaloriesPerMeal(bmi, bmr, desiredCalories)

	// Check if calories per meal is valid
	if caloriesPerMeal <= 0 {
		c.JSON(http.StatusBadRequest, "Invalid calories per meal value")
		return
	}

	// Generate meal suggestions based on calculated calories per meal
	mealSuggestions := generateMealSuggestions(caloriesPerMeal)

	// Return meal suggestions in JSON format
	c.JSON(http.StatusOK, mealSuggestions)
}
