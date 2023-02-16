package middelware

import (
	"fmt"
	"gin/controller/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(auth.JwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Token is expired"})
			panic("invalid token")

		} else {
			c.JSON(http.StatusAccepted, gin.H{"message": "Token is valid"})
			claims := token.Claims.(jwt.MapClaims)["id"]
			claim := token.Claims.(jwt.MapClaims)["role"]
			email := token.Claims.(jwt.MapClaims)["email"]
			fmt.Println("token vaild")
			c.Set("id", claims)
			c.Set("role", claim)
			c.Set("email", email)
			c.Next()

		}

	}
}
