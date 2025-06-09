package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

		var tokenString string

		if strings.HasPrefix(authHeader,"Bearer ") {
			tokenString = strings.TrimPrefix(authHeader,"Bearer ")
		} else {
			cookieToken, err := c.Cookie("jwt")
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Missing token at the header or cookie"})
				c.Abort()
				return
			}
			tokenString = cookieToken
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface {}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})


		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{ "error": "Invalid or token has expired"});
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, exists := claims["userId"]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing userId in token claims"})
			c.Abort()
			return
		}

		c.Set("userId", userID)
	}


		c.Next()
	}
}