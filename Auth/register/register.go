package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	entity "e-Commerce/Models/User_Model"
	"e-Commerce/database"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	var err entity.Error
	var users = []entity.User{}
	// here we want to decode the user
	json.NewDecoder(r.Body).Decode(&user)
	// here we want to fetch the data from database as user input
	database.Database.Where("email = ?", user.Email).Find(&users)
	// here we can give the condition.
	if len(users) == 0 {

		password := []byte(string(user.Password))
		// here we can create hashed Passwword by bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
		if err != nil {
			panic(err)
		}
		// Comparing the password with the hash
		err = bcrypt.CompareHashAndPassword(hashedPassword, password)
		fmt.Println(err) // nil means it is a match

		// here we can convert hashed password in string form and store in user Password
		user.Password = string(hashedPassword)

		// // here we can create the data on the database but the password will be saved in hashed Password
		database.Database.Create(&user)
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusCreated)

		// fmt.Println(password)
		// here we want to send email message on user email.
		// smtps.Smtp(user.Email, "Subject: Registration \r\n\r\n"+"Hi, "+user.First_Name+user.Last_Name+" \nYour Account has been created successfully"+
		// 	"\n Email : "+user.Email+"\n Password : "+string(password))
		// // fmt.Println(user.Phone)
		// // here we want to send a sms message on user phone number
		// sms.SMS(user.Phone, "Hello, "+user.Username+" Your Account has been created successfully")

	} else {
		err = entity.Error{Message: "This Email already exist"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
	}

}
