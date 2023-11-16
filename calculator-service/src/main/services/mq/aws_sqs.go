package services

import (
	"log"

	"github.com/mookrob/servicecalculator/main/models"
	"github.com/mookrob/servicecalculator/main/repositories"
)

type UserService struct {
	MongoDBRepo *repositories.MongoDBRepository
	SQSRepo     *repositories.SQSRepository
}

func NewUserService(mongoDBRepo *repositories.MongoDBRepository, sqsRepo *repositories.SQSRepository) *UserService {
	return &UserService{
		MongoDBRepo: mongoDBRepo,
		SQSRepo:     sqsRepo,
	}
}

func (us *UserService) SaveUserBodyData(userData *models.UserBodyData) error {
	err := us.MongoDBRepo.InsertUserBodyData(userData)
	if err != nil {
		log.Printf("Failed to save to MongoDB: %v", err)
		return err
	}

	err = us.SQSRepo.SendMessage(userData)
	if err != nil {
		log.Printf("Failed to send message to SQS: %v", err)
		return err
	}

	return nil
}
