package grpc_services

import (
	"log"

	"github.com/google/uuid"
	pb "github.com/mookrob/servicemeal/main/grpc-server"
	"github.com/mookrob/servicemeal/main/models"
	repositories "github.com/mookrob/servicemeal/main/repositories"
)

type MealServer struct {
	UserMealRepository *repositories.UserMealRepository
}

func NewUserMealGrpcService(r *repositories.UserMealRepository) *MealServer {
	return &MealServer{UserMealRepository: r}
}

func (s *MealServer) mustEmbedUnimplementedMealServer() {}

func (s *MealServer) GetUserFavMeal(req *pb.GetUserFavMealRequest, stream pb.Meal_GetUserFavMealServer) error {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed parse param: %v", err)
		return err
	}

	// call repo
	rows, err := s.UserMealRepository.GetUserFavMealByUserId(id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed on user meal repository call: %v", err)
		return err
	}
	defer rows.Close()

	// iterate through the result set and scan each row into a MealItem model and response
	for rows.Next() {
		var meal models.Meal
		if err := rows.Scan(&meal.Id, &meal.Name, &meal.Energy, &meal.Protein, &meal.Carbohydrate, &meal.Fat,
			&meal.Sodium, &meal.Cholesterol, &meal.CreatedAt, &meal.UpdatedAt); err != nil {
			log.Printf("grpc GetUserFavMeal failed on row scan: %v", err)
			return err
		}

		// TODO:: send nil value through grpc // current replace nil with 0
		if meal.Protein == nil {
			meal.Protein = new(float64)
		}
		if meal.Carbohydrate == nil {
			meal.Carbohydrate = new(float64)
		}
		if meal.Fat == nil {
			meal.Fat = new(float64)
		}
		if meal.Sodium == nil {
			meal.Sodium = new(float64)
		}
		if meal.Cholesterol == nil {
			meal.Cholesterol = new(float64)
		}

		// build response
		response := &pb.MealItem{
			Id:           meal.Id.String(),
			Name:         meal.Name,
			Energy:       meal.Energy,
			Protein:      *meal.Protein,
			Carbohydrate: *meal.Carbohydrate,
			Fat:          *meal.Fat,
			Sodium:       *meal.Sodium,
			Cholesterol:  *meal.Cholesterol,
		}

		// send grpc stream
		if err := stream.Send(response); err != nil {
			log.Printf("grpc GetUserFavMeal failed on row scan: %v", err)
			return err
		}

	}

	if err := rows.Err(); err != nil {
		log.Printf("grpc GetUserFavMeal failed query rows: %v", err)
		return err
	}

	return nil
}
