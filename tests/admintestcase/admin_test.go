package admintestcase

import (
	"fmt"
	"gin/controller/auth"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Test_AuthRequired() gin.HandlerFunc {
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

func Test_Post_Admin_Store(t *testing.T) {
	r := gin.Default()
	r.POST("/fixstore", Test_AuthRequired(), func(c *gin.Context) {

		// Auth := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNvaGFuQGdtYWlsLmNvbSIsImV4cCI6MTY3Njg3NjU4MiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.4n3NBF0ncSzsMqqgIPGhZvMRlBtEOG6MXq9LdBOjCg4"

		var Store struct {
			Store_Name    string `json:"store_name"`
			Store_Address string `json:"store_address"`
			Vendor_Id     int    `json:"vendor_id"`
		}
		if err := c.BindJSON(&Store); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		c.JSON(http.StatusOK, Store)
	})

	req, _ := http.NewRequest("POST", "/fixstore", strings.NewReader(`{"store_name":"Jigyasa Store",
	"store_address":"shri Radhe Lal Market,sector-135,",
	"vendor_id":"3"}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, resp.Code)
	}

	expectedBody := `{"store_name":"Jigyasa Store",
	"store_address":"shri Radhe Lal Market,sector-135,",
	"vendor_id":"3"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("expected response body %v but got %v", expectedBody, resp.Body.String())
		// fmt.Println("")
	}
}
