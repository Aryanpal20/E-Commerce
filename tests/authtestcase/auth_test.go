package admintestcase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin/controller/auth"
	entity "gin/models/user_model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// var DB *gorm.DB

// func ConnectDatabase() {
// 	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {"gorm.io/driver/sqlite"

// 		panic("Failed to connect to database!")
// 	}
// 	database.AutoMigrate(entity.User{})
// 	DB = database
// }

func Test_Register(t *testing.T) {
	// ConnectDatabase()
	r := gin.Default()
	var user entity.User
	r.POST("/register", auth.Register)
	user = entity.User{
		Email:      "sohan@gmail.com",
		Password:   "sohan",
		First_Name: "Sohan",
		Last_Name:  "Kumar",
		Role:       "admin",
		CreatedAT:  time.Now(),
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// r.POST("/register", func(c *gin.Context) {

	// 	// Parse the request body to get the user data
	// 	user = entity.User{
	// 	Email:      "sohan@gmail.com",
	// 	Password:   "sohan",
	// 	First_Name: "Sohan",
	// 	Last_Name:  "Kumar",
	// 	Role:       "admin",
	// 	CreatedAT:  time.Now(),
	// }

	// 	if err := c.BindJSON(&user); err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	// 		return
	// 	}

	// 	// Create the user in the database
	// 	// ...

	// 	c.JSON(http.StatusOK, user)
	// })
	// // Create a test request with valid user data in the requesgin.SetMode(gin.ReleaseMode)t body
	// req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"email":"sohan@gmail.com","password":"sohan","first_name":"Sohan",
	// "last_name":"Kumar","role":"admin"}`))
	// req.Header.Set("Content-Type", "application/json")
	// resp := httptest.NewRecorder()

	// // Perform the request on the test router
	// r.ServeHTTP(resp, req)

	// // Assert that the response code is 200 OK
	// if resp.Code != http.StatusOK {
	// 	t.Errorf("expected status code %v but got %v", http.StatusOK, resp.Code)
	// }

	// // Assert that the response body contains the correct user data
	// expectedBody := `{"email":"sohan@gmail.com","password":"sohan","first_name":"Sohan","last_name":"Kumar","role":"admin" }`
	// if resp.Body.String() != expectedBody {
	// 	t.Errorf("expected response body %v but got %v", expectedBody, resp.Body.String())
	// 	fmt.Println(resp.Body, "vlkhvusdhjls")
	// 	fmt.Println("lvksndjnbb", expectedBody)
	// 	return
	// }
	// // fmt.Println("efjhsdjkvhjv", resp.Body)
	// DB.Create(&user)
	// fmt.Println("kdvbjbvdkv")

}
func Test_Wrong_Email(t *testing.T) {
	// ConnectDatabase()
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var user struct {
			Email      string `json:"email"`
			Password   string `json:"password"`
			First_Name string `json:"first_name"`
			Last_Name  string `json:"last_name"`
			Role       string `json:"role"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"email":"sohangmail.com","password":"sohan","first_name":"Sohan",
	"last_name":"Kumar","role":"admin"}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, resp.Code)
	}

	expectedBody := `{"email":"sohan@gmail.com","password":" gin.SetMode(gin.ReleaseMode)sohan","first_name":"Sohan","last_name":"Kumar","role":"admin"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("expected response body %v but got %v", expectedBody, resp.Body.String())
		fmt.Println("Enter Right Email Format")
		fmt.Println("")
		fmt.Println("klvnrjkbvnjkerngj nre iurwghvrrei nugigf bhfiuvhjmcdlfk,x", resp.Body.String())
		return
	}
	// DB.Create()
}

func Test_Login(t *testing.T) {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var user struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"email":"sohan@gmail.com","password":"sohan"}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, resp.Code)
	}

	expectedBody := `{"email":"sohan@gmail.com","password":"sohan"}`
	fmt.Println(resp.Body.String())
	if resp.Body.String() != expectedBody {
		t.Errorf("expected response body %v but got %v", expectedBody, resp.Body.String())
		fmt.Println("Enter Email Format")
	}
}

func Test_Wrong_Password(t *testing.T) {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var user struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"email":"sohan@gmail.com","password":"sohan"}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, resp.Code)
	}

	expectedBody := `{"email":"sohan@gmail.com","password":"sohan12"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("expected response body %v but got %v", expectedBody, resp.Body.String())
		fmt.Println("E-Mail or Password is incorrect")
	}
}
func Test_Wrong_Login_Email(t *testing.T) {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var user struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"email":"sohan12@gmail.com","password":"sohan"}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, resp.Code)
	}

	expectedBody := `{"email":"sohan@gmail.com","password":"sohan"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("expected response body %v but got %v", expectedBody, resp.Body.String())
		fmt.Println("E-Mail or Password is incorrect")
	}
}
