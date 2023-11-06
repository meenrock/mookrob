package services

import (
	"context"
	"io"
	"log"
	"net/http"

	enums "github.com/mookrob/serviceuser/main/enums"
	pb "github.com/mookrob/serviceuser/main/grpc-client/meal"
	repositories "github.com/mookrob/serviceuser/main/repositories"
	utils "github.com/mookrob/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserRestService struct {
	UserRepository *repositories.UserRepository
	mealGrpcHost   string
}

func NewUserRestService(r *repositories.UserRepository) *UserRestService {

	MEAL_GRPC_HOST := viper.GetString("client.meal-grpc-host")

	return &UserRestService{UserRepository: r, mealGrpcHost: MEAL_GRPC_HOST}
}

// GetUserById response model
type UserDetailResponse struct {
	Id          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Status      enums.Status `json:"status"`
	PhoneNumber *string      `json:"phone_number"`
	Email       string       `json:"email"`
}

func (s *UserRestService) GetUserById(ctx *gin.Context) {
	userDataRaw, exist := ctx.Get("userData")
	if exist != true {
		log.Printf("rest GetUserFavFoodByUserId failed parse userData")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	userData, ok := utils.ExtractUserData(userDataRaw)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	// call repo
	user, err := s.UserRepository.GetUserById(userData.UserId)
	if err != nil {
		log.Printf("GetUserById failed on user repository call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	// build response
	userResponseDto := UserDetailResponse{
		Id:          user.Id,
		Name:        user.NickName,
		Status:      user.Status,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}

	ctx.JSON(http.StatusOK, userResponseDto)
}

// GetUserFavFoodByUserId response model
type UserFavFoodResponse struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Energy       float64  `json:"energy"`
	Protein      *float64 `json:"protein"`
	Carbohydrate *float64 `json:"carbohydrate"`
	Fat          *float64 `json:"fat"`
	Sodium       *float64 `json:"sodium"`
	Cholesterol  *float64 `json:"cholesterol"`
}

func (s *UserRestService) GetUserFavFoodByUserId(ctx *gin.Context) {
	userDataRaw, exist := ctx.Get("userData")
	if exist != true {
		log.Printf("rest GetUserFavFoodByUserId failed parse userData")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	userData, ok := utils.ExtractUserData(userDataRaw)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	// connect meal service
	conn, err := grpc.Dial(s.mealGrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("rest GetUserFavFoodByUserId failed to connect meal server: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer conn.Close()

	// define meal client
	mealCon := pb.NewMealClient(conn)

	// build GetUserFavFood request
	req := &pb.GetUserFavFoodRequest{
		Id: userData.UserId.String(),
	}

	// send grpc request
	stream, err := mealCon.GetUserFavFood(context.Background(), req)
	if err != nil {
		log.Printf("rest GetUserFavFoodByUserId failed RPC: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	var userFavResponse []UserFavFoodResponse
	// loop recieve stream data
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// end of the stream, no more data
			break
		}
		if err != nil {
			log.Printf("rest GetUserFavFoodByUserId failed RPC meal stream: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		// build response
		userFav := UserFavFoodResponse{
			Id:           response.Id,
			Name:         response.Name,
			Energy:       response.Energy,
			Protein:      &response.Protein,
			Carbohydrate: &response.Carbohydrate,
			Fat:          &response.Fat,
			Sodium:       &response.Sodium,
			Cholesterol:  &response.Cholesterol,
		}
		userFavResponse = append(userFavResponse, userFav)
	}

	ctx.JSON(http.StatusOK, userFavResponse)
}
