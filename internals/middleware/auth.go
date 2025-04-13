package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Authmiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authoriz := c.GetHeader("Authorization")
		if len(authoriz) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "please provide Auth token",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authoriz, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "unexpected signin method",
				})

				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

			}
			return []byte("your_secret_key"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Pass user data to next handler
			log.Println("email", claims["email"])
			c.Set("email", claims["email"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT claims"})
			c.Abort()
		}
	}

}
