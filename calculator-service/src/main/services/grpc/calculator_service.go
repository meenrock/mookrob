package services

import (
	"log"

	"github.com/google/uuid"
	pb "github.com/mookrob/servicecalculator/main/grpc-server"
	"github.com/mookrob/servicecalculator/main/models"
	repositories "github.com/mookrob/servicecalculator/main/repositories"
)

type CalculatorServer struct {
	UserCalculatorRepository *repositories.UserCalculatorRepository
}

func NewCalculatorGrpcService(r *repositories.UserCalculatorRepository) *CalculatorServer {
	return &CalculatorServer{UserCalculatorRepository: r}
}

func (s *CalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (s *CalculatorServer) GetBMI(req *pb.GetUserBMIRequest, stream pb.BMIResponse) (*pb.BMIResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed parse param: %v", err)
		return &pb.BMIResponse{}, err
	}
	// call repo
	collection, err := s.UserCalculatorRepository.GetUserCalculationByUserId(id)
	if err != nil {
		log.Printf("grpc GetUserFavMeal failed on user meal repository call: %v", err)
		return &pb.BMIResponse{}, err
	}

	var parameter models.UserBodyData

	return &pb.BMIResponse{
		Id:  parameter.ID.String(),
		Bmi: collection.InsertedID,
	}, nil

}
