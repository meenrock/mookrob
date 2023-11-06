package rest_services

import (
	"log"
	"net/http"
	"time"

	"github.com/mookrob/serviceauth/main/models"
	repositories "github.com/mookrob/serviceauth/main/repositories"
	constants "github.com/mookrob/shared/constants"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte
var refreshKey []byte

type AuthenticationRestService struct {
	AuthenticationRepository *repositories.AuthenticationRepository
}

func NewAuthenticationRestService(r *repositories.AuthenticationRepository) *AuthenticationRestService {
	jwtKey = []byte(viper.GetString("jwt.secret_key"))
	refreshKey = []byte(viper.GetString("jwt.refresh_key"))
	return &AuthenticationRestService{AuthenticationRepository: r}
}

// Login request model
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login response model
type LoginResponse struct {
	Role         constants.Role `json:"role"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
}

func (s *AuthenticationRestService) Login(ctx *gin.Context) {

	var request LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call repo
	auth, err := s.AuthenticationRepository.GetAuthenticationByUsernameAndStatusActive(request.Username)
	if err != nil {
		log.Println("Login: Failed on user repository call: ", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "username not found."})
		return
	}

	err = verifyPassword(request.Password, auth.Password)
	if err != nil {
		log.Println("Login: Password verification failed:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password."})
		return
	}

	token, refreshToken, err := generateJWT(auth)
	if err != nil {
		log.Println("Login: Generate token failed:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal."})
		return
	}

	// build response
	loginResponseDto := LoginResponse{
		Role:         auth.Role,
		Token:        token,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, loginResponseDto)
}

// Function to hash the password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Function to verify the password
func verifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func generateJWT(auth models.AuthenticationData) (string, string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["role"] = auth.Role
	claims["user_name"] = auth.Username
	claims["user_id"] = auth.UserId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Create token with claims and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	claims["token_type"] = auth.Role
	claims["user_name"] = auth.Username
	claims["user_id"] = auth.UserId
	refreshClaims["exp"] = time.Now().Add(time.Hour * 36).Unix() // Refresh token expires in 36 hours
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}
