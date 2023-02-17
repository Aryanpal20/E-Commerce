package auth

import (
	"fmt"
	"gin/database"
	entity "gin/models/user_model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtToken struct {
	Token string `json:"token"`
}

var JwtKey = []byte("Jwt_Key")

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var users entity.User
	database.Database.Where("email = ?", email).First(&users)
	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	fmt.Println(err)

	if err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    users.ID,
			"email": users.Email,
			"role":  users.Role,
			"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		})
		tokenString, error := token.SignedString(JwtKey)
		c.JSON(http.StatusAccepted, gin.H{"Token": tokenString})
		if error != nil {
			fmt.Println(error)
		}
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "E-Mail or Password is incorrect"})
	}
}

func Register(c *gin.Context) {

	var input entity.User
	var users entity.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := entity.User{Email: input.Email, Password: input.Password, First_Name: input.First_Name, Last_Name: input.Last_Name, Role: input.Role, CreatedAT: time.Now()}

	if !strings.Contains(user.Email, "@") {
		c.JSON(400, gin.H{"MESSAGE": "Enter email Format"})
		return
	}
	database.Database.Where("email = ?", user.Email).First(&users)

	if users.Email != user.Email {
		password := []byte(string(user.Password))
		hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
		if err != nil {
			panic(err)
		}
		err = bcrypt.CompareHashAndPassword(hashedPassword, password)
		fmt.Println(err)
		user.Password = string(hashedPassword)
		fmt.Println("ABCD", user.Password)

		// database.Database.Create(&user)
		c.JSON(http.StatusOK, gin.H{"Data": "registration successfully", "Msg": user})
	} else {
		fmt.Println("djvbdshvbhk")
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email already exist"})
	}
}
