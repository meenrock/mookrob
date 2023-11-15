package services

import (
	"log"

	"github.com/your-username/myproject/model"
	"github.com/your-username/myproject/repository"
)

type UserService struct {
	MongoDBRepo *repository.MongoDBRepository
	SQSRepo     *repository.SQSRepository
}

func NewUserService(mongoDBRepo *repository.MongoDBRepository, sqsRepo *repository.SQSRepository) *UserService {
	return &UserService{
		MongoDBRepo: mongoDBRepo,
		SQSRepo:     sqsRepo,
	}
}

func (us *UserService) SaveUserBodyData(userData *model.UserBodyData) error {
	// Save to MongoDB
	err := us.MongoDBRepo.InsertUserBodyData(userData)
	if err != nil {
		log.Printf("Failed to save to MongoDB: %v", err)
		return err
	}

	// Send to SQS
	err = us.SQSRepo.SendMessage(userData)
	if err != nil {
		log.Printf("Failed to send message to SQS: %v", err)
		return err
	}

	return nil
}
