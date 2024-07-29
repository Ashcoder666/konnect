package middlewares

import (
	"fmt"
	"net/http"

	"github.com/ashcoder666/konnect/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 1"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		// fmt.Println(tokenString)
		// Parse JWT token
		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		fmt.Println(*token)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 2"})
			c.Abort()
			return
		}

		// Add claims to context
		c.Set("email", claims.Email)

		c.Next()
	}
}
