package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	pb "github.com/mookrob/servicecalculator/main/grpc-server"
	"github.com/mookrob/servicecalculator/main/models"
	"github.com/mookrob/servicecalculator/main/repositories"
)

type CalculatorServer struct {
	CalculatorUserRepository *repositories.UserCalculatorRepositoryMongo
}

func NewCalculatorGrpcService(r *repositories.UserCalculatorRepositoryMongo) *CalculatorServer {
	return &CalculatorServer{CalculatorUserRepository: r}
}

func (s *CalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (s *CalculatorServer) GetBMI(input context.Context, req *pb.GetUserBMIRequest) (*pb.BMIResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed parse param: %v", err)
		return &pb.BMIResponse{}, err
	}
	// call repo
	collection, err := s.CalculatorUserRepository.GetUserCalculationBMI(id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed on user meal repository call: %v", err)
		return &pb.BMIResponse{}, err
	}

	// Decode the result into the UserBodyData struct
	var parameter models.UserBodyData
	if err := collection.Decode(&parameter); err != nil {
		log.Printf("Error decoding UserBodyData: %v", err)
		return &pb.BMIResponse{}, err
	}

	return &pb.BMIResponse{
		Id:  parameter.ID.String(),
		Bmi: parameter.BMI,
	}, nil

}

func (s *CalculatorServer) GetBMR(input context.Context, req *pb.GetUserBMRRequest) (*pb.BMRResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed parse param: %v", err)
		return &pb.BMRResponse{}, err
	}
	// call repo
	collection, err := s.CalculatorUserRepository.GetUserCalculationBMR(id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed on user meal repository call: %v", err)
		return &pb.BMRResponse{}, err
	}

	// Decode the result into the UserBodyData struct
	var parameter models.UserBodyData
	if err := collection.Decode(&parameter); err != nil {
		log.Printf("Error decoding UserBodyData: %v", err)
		return &pb.BMRResponse{}, err
	}

	return &pb.BMRResponse{
		Id:  parameter.ID.String(),
		Bmr: parameter.BMR,
	}, nil

}
