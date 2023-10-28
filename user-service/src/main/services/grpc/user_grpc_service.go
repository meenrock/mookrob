package service

import (
	"log"

	"github.com/google/uuid"
	pb "github.com/mookrob/serviceuser/main/grpc-server/github.com/mookrob"
	repositories "github.com/mookrob/serviceuser/main/repositories"
)

type UserServer struct {
	UserRepository *repositories.UserRepository
}

func NewUserGrpcService(r *repositories.UserRepository) *UserServer {
	return &UserServer{UserRepository: r}
}

func (s *UserServer) mustEmbedUnimplementedUserServer() {}

func (s *UserServer) EditUserByUserId(req *pb.GetUserIdRequest, stream pb.User_EditUserByIdServer) error {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("grpc GetUserFavFood failed parse param: %v", err)
		return err
	}

	// call repo
	row, err := s.UserRepository.EditUserById(id)
	if err != nil {
		log.Printf("grpc GetUserFavFood failed on user food repository call: %v", err)
		return err
	}
	// defer row.Close()

	// // iterate through the result set and scan each row into a FoodItem model and response
	// for rows.Next() {
	// 	var user models.User
	// 	if err := rows.Scan(&user.Id, &user.Status, &user.FirstName, &user.LastName, &user.NickName, &user.PhoneNumber,
	// 		&user.Email, &user.Gender, &user.Age, &user.Height, &user.Weight,
	// 		&user.ExpectedBmi,&user.CreatedAt, &user.UpdatedAt); err != nil {
	// 		log.Printf("grpc GetUserFavFood failed on row scan: %v", err)
	// 		return err
	// 	}

	// 	// TODO:: send nil value through grpc // current replace nil with 0
	// 	if food.Protein == nil {
	// 		food.Protein = new(float64)
	// 	}
	// 	if food.Carbohydrate == nil {
	// 		food.Carbohydrate = new(float64)
	// 	}
	// 	if food.Fat == nil {
	// 		food.Fat = new(float64)
	// 	}
	// 	if food.Sodium == nil {
	// 		food.Sodium = new(float64)
	// 	}
	// 	if food.Cholesterol == nil {
	// 		food.Cholesterol = new(float64)
	// 	}

	// build response
	response := &pb.UserItem{
		Id:          row.Id.String(),
		Firstname:   row.FirstName,
		Lastname:    row.LastName,
		Nickname:    row.NickName,
		Phonenumber: *row.PhoneNumber,
		Email:       row.Email,
		Age:         float64(row.Age),
		Height:      row.Height,
		Weight:      row.Weight,
		Expectedbmi: *row.ExpectedBmi,
	}

	// send grpc stream
	if err := stream.Send(response); err != nil {
		log.Printf("grpc GetUserFavFood failed on row scan: %v", err)
		return err
	}

	return nil
}
