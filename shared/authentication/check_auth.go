package authentication

import (
	"log"
	"net/http"
	"slices"
	"strings"

	constants "github.com/mookrob/shared/constants"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware(role constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtKey = []byte(viper.GetString("JWT_SECRET_KEY"))
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			log.Println("AuthMiddleware: Invalid token", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("AuthMiddleware: Failed to parse access claims")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse access token"})
			return
		}

		if string(role) != claims["role"] {
			log.Println("AuthMiddleware: Invalid role")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unauthenticated"})
			return
		}

		c.Set("userData", claims)

		c.Next()
	}
}

func AuthListMiddleware(roleList []constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtKey = []byte(viper.GetString("JWT_SECRET_KEY"))
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			log.Println("AuthMiddleware: Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("AuthMiddleware: Failed to parse access claims")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse access token"})
			return
		}

		var userRole constants.Role
		ok = userRole.ParseString(claims["role"].(string))
		if !ok {
			log.Println("AuthMiddleware: Failed to parse role claims")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse role from token"})
			return
		}

		if !slices.Contains(roleList, userRole) {
			log.Println("AuthMiddleware: Invalid role: ", userRole)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unauthenticated"})
			return
		}

		c.Set("userData", claims)

		c.Next()
	}
}
