package customerdelete

import (
	email "e-Commerce/FetchOnlyOneData/fetchemail"
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	cust "e-Commerce/Models/Customer_Model"
	users "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CustomerDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	email := email.Task_creator(token)
	role := role.Is_manager(token)
	if role == "customer" {
		var err users.Error
		var user users.User
		database.Database.First(&user, mux.Vars(r)["customer_id"])
		fmt.Println(user)
		if email == user.Email {
			var cust cust.Customer
			database.Database.Where("customer_id = ?", user.ID).Delete(&cust)
			fmt.Println(cust)
			err = users.Error{Message: "your order will be deleted"}
			json.NewEncoder(w).Encode(err)
			w.WriteHeader(http.StatusOK)
		} else {
			err = users.Error{Message: "you can't delete another's data"}
			json.NewEncoder(w).Encode(err)
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}
