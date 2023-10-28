package grpc_services

import (
	"context"
	"log"

	pb "github.com/mookrob/serviceauth/main/grpc-server"
	models "github.com/mookrob/serviceauth/main/models"
	repositories "github.com/mookrob/serviceauth/main/repositories"
	constants "github.com/mookrob/shared/constants"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServer struct {
	AuthenticationRepository *repositories.AuthenticationRepository
}

func NewAuthenticationGrpcService(r *repositories.AuthenticationRepository) *AuthServer {
	return &AuthServer{AuthenticationRepository: r}
}

func (s *AuthServer) mustEmbedUnimplementedAuthServer() {}

func (s *AuthServer) CreateAuthUser(ctx context.Context, req *pb.CreateAuthUserRequest) (*pb.CreateAuthUserResponse, error) {

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("CreateAuthUser: Error on encrypt: ", err)
		return &pb.CreateAuthUserResponse{}, err
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		log.Println("CreateAuthUser: Error on parse user id: ", err)
		return &pb.CreateAuthUserResponse{}, err
	}

	newAuthData := models.AuthenticationData{
		Username: req.Username,
		Password: string(hashedPassword),
		UserId:   &userId,
	}
	// call repo
	id, err := s.AuthenticationRepository.CreateAuthenticationUser(newAuthData)
	if err != nil {
		log.Println("CreateAuthUser: Failed on user repository call: ", err)
		return &pb.CreateAuthUserResponse{}, err
	}
	log.Println("CreateAuthUser: Create auth user succeed: ", id)

	return &pb.CreateAuthUserResponse{
		Role:     string(constants.GENERAL_USER),
		Username: req.Username,
	}, nil
}
