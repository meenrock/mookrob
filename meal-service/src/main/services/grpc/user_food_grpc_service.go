package grpc_services

import (
	"log"

	"github.com/google/uuid"
	pb "github.com/mookrob/servicemeal/main/grpc-server"
	"github.com/mookrob/servicemeal/main/models"
	repositories "github.com/mookrob/servicemeal/main/repositories"
)

type MealServer struct {
	UserFoodRepository *repositories.UserFoodRepository
}

func NewUserFoodGrpcService(r *repositories.UserFoodRepository) *MealServer {
	return &MealServer{UserFoodRepository: r}
}

func (s *MealServer) mustEmbedUnimplementedMealServer() {}

func (s *MealServer) GetUserFavFood(req *pb.GetUserFavFoodRequest, stream pb.Meal_GetUserFavFoodServer) error {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("grpc GetUserFavFood failed parse param: %v", err)
		return err
	}

	// call repo
	rows, err := s.UserFoodRepository.GetUserFavFoodByUserId(id)
	if err != nil {
		log.Printf("grpc GetUserFavFood failed on user food repository call: %v", err)
		return err
	}
	defer rows.Close()

	// iterate through the result set and scan each row into a FoodItem model and response
	for rows.Next() {
		var food models.Food
		if err := rows.Scan(&food.Id, &food.Name, &food.Energy, &food.Protein, &food.Carbohydrate, &food.Fat,
			&food.Sodium, &food.Cholesterol, &food.CreatedAt, &food.UpdatedAt); err != nil {
			log.Printf("grpc GetUserFavFood failed on row scan: %v", err)
			return err
		}

		// TODO:: send nil value through grpc // current replace nil with 0
		if food.Protein == nil {
			food.Protein = new(float64)
		}
		if food.Carbohydrate == nil {
			food.Carbohydrate = new(float64)
		}
		if food.Fat == nil {
			food.Fat = new(float64)
		}
		if food.Sodium == nil {
			food.Sodium = new(float64)
		}
		if food.Cholesterol == nil {
			food.Cholesterol = new(float64)
		}

		// build response
		response := &pb.FoodItem{
			Id:           food.Id.String(),
			Name:         food.Name,
			Energy:       food.Energy,
			Protein:      *food.Protein,
			Carbohydrate: *food.Carbohydrate,
			Fat:          *food.Fat,
			Sodium:       *food.Sodium,
			Cholesterol:  *food.Cholesterol,
		}

		// send grpc stream
		if err := stream.Send(response); err != nil {
			log.Printf("grpc GetUserFavFood failed on row scan: %v", err)
			return err
		}

	}

	if err := rows.Err(); err != nil {
		log.Printf("grpc GetUserFavFood failed query rows: %v", err)
		return err
	}

	return nil
}
