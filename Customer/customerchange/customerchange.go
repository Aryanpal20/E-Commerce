package customerchange

import (
	email "e-Commerce/FetchOnlyOneData/fetchemail"
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	cust "e-Commerce/Models/Customer_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CustomerChange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	if role == "customer" {
		email := email.Task_creator(token)
		var cust cust.Customer
		var user user.User
		json.NewDecoder(r.Body).Decode(&cust)
		database.Database.Where("id = ?", cust.Customer_Id).Find(&user)
		if email == user.Email {
			if err := database.Database.Model(&cust).Where("customer_id = ?", cust.Customer_Id).Update("select_product_quantity", cust.Select_Product_Quantity).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			database.Database.Save(cust)
			json.NewEncoder(w).Encode(cust)
		} else {
			email := "Please enter your I'D"
			json.NewEncoder(w).Encode(email)
		}
	} else {
		role := "You are not the Customer"
		json.NewEncoder(w).Encode(role)
	}
}
