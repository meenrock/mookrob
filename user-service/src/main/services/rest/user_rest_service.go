package services

import (
	"context"
	"io"
	"log"
	"net/http"

	enums "github.com/mookrob/serviceuser/main/enums"
	pb_auth "github.com/mookrob/serviceuser/main/grpc-client/auth"
	pb_meal "github.com/mookrob/serviceuser/main/grpc-client/meal"
	models "github.com/mookrob/serviceuser/main/models"
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
	authGrpcHost   string
}

func NewUserRestService(r *repositories.UserRepository) *UserRestService {

	MEAL_GRPC_HOST := viper.GetString("CLIENT_MEAL_GRPC_HOST")
	AUTH_GRPC_HOST := viper.GetString("CLIENT_AUTH_GRPC_HOST")

	return &UserRestService{UserRepository: r, mealGrpcHost: MEAL_GRPC_HOST, authGrpcHost: AUTH_GRPC_HOST}
}

// CreateUser request model
type CreateUserRequest struct {
	FirstName   string   `json:"first_name" binding:"required"`
	LastName    string   `json:"last_name" binding:"required"`
	NickName    string   `json:"nick_name" binding:"required"`
	PhoneNumber *string  `json:"phone_number"`
	Email       string   `json:"email" binding:"required,email"`
	Gender      string   `json:"gender" binding:"required"`
	Age         int64    `json:"age" binding:"required"`
	Height      float64  `json:"height" binding:"required"`
	Weight      float64  `json:"weight" binding:"required"`
	ExpectedBmi *float64 `json:"expected_bmi"`
	Password    string   `json:"password" binding:"required"`
}

func (s *UserRestService) CreateUser(ctx *gin.Context) {

	var request CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("rest CreateUser: error on parse request: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist, err := s.UserRepository.ExistByEmail(request.Email)
	if err != nil {
		log.Println("rest CreateUser: error on query exist by mail: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	if exist == true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exist"})
		return
	}

	conn, err := grpc.Dial(s.authGrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("rest CreateUser: failed to connect auth server: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer conn.Close()

	// define auth client
	authCon := pb_auth.NewAuthClient(conn)

	// create user
	newUser := models.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		NickName:    request.NickName,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Gender:      request.Gender,
		Age:         request.Age,
		Height:      request.Height,
		Weight:      request.Weight,
		ExpectedBmi: request.ExpectedBmi,
	}
	newUserId, err := s.UserRepository.CreateUser(newUser)
	if err != nil {
		log.Println("rest GetUserById: failed on user repository call: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	// build CreateAuthUser request
	authReq := &pb_auth.CreateAuthUserRequest{
		Username: request.Email,
		Password: request.Password,
		UserId:   newUserId.String(),
	}

	// create auth user
	authRes, err := authCon.CreateAuthUser(context.Background(), authReq)
	if err != nil {
		log.Println("rest CreateUser: could not create authentication user: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	// TODO:: calculate bmi

	ctx.JSON(http.StatusOK, gin.H{"role": authRes.Role, "username": authRes.Username})
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
		log.Println("rest GetUserById: failed parse userData")
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
		log.Println("rest GetUserById: failed on user repository call: ", err)
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

// GetUserFavMealByUserId response model
type UserFavMealResponse struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Energy       float64  `json:"energy"`
	Protein      *float64 `json:"protein"`
	Carbohydrate *float64 `json:"carbohydrate"`
	Fat          *float64 `json:"fat"`
	Sodium       *float64 `json:"sodium"`
	Cholesterol  *float64 `json:"cholesterol"`
}

func (s *UserRestService) GetUserFavMealByUserId(ctx *gin.Context) {
	userDataRaw, exist := ctx.Get("userData")
	if exist != true {
		log.Println("rest GetUserFavMealByUserId: failed parse userData")
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
		log.Println("rest GetUserFavMealByUserId: failed to connect meal server: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}
	defer conn.Close()

	// define meal client
	mealCon := pb_meal.NewMealClient(conn)

	// build GetUserFavMeal request
	req := &pb_meal.GetUserFavMealRequest{
		Id: userData.UserId.String(),
	}

	// send grpc request
	stream, err := mealCon.GetUserFavMeal(context.Background(), req)
	if err != nil {
		log.Println("rest GetUserFavMealByUserId: failed RPC: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
		return
	}

	var userFavResponse []UserFavMealResponse
	// loop recieve stream data
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// end of the stream, no more data
			break
		}
		if err != nil {
			log.Println("rest GetUserFavMealByUserId: failed RPC meal stream: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
			return
		}

		// build response
		userFav := UserFavMealResponse{
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
