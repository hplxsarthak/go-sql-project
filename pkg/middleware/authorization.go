package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")

		_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
		}
	}
}
