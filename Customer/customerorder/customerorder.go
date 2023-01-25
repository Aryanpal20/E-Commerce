package customerorder

import (
	email "e-Commerce/FetchOnlyOneData/fetchemail"
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	"fmt"

	customer "e-Commerce/Models/Customer_Model"
	orders "e-Commerce/Models/Order_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"net/http"
	"strings"
)

func Customer_Order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	fmt.Println(role)
	email := email.Task_creator(token)
	fmt.Println(email)
	var err user.Error
	if role == "customer" {
		var user user.User
		var cust customer.Customer
		var order orders.Order
		json.NewDecoder(r.Body).Decode(&order)
		database.Database.Where("id = ?", order.OrderID).Find(&user)
		fmt.Println(user)
		if email == user.Email {
			database.Database.Where("customer_id = ?", order.OrderID).Find(&cust)
			fmt.Println(cust)
			// order.Userid = user.ID
			order.Product_Id = cust.Product_Id
			order.Product_Name = cust.Select_Product_Name
			order.Product_Quantity = cust.Select_Product_Quantity
			order.Total_Rate = cust.Quantity_wise_Rate
			database.Database.Create(&order)
			json.NewEncoder(w).Encode(&order)
		}
	} else {
		err = user.Error{Message: "You are not a Customer"}
		json.NewEncoder(w).Encode(err)
	}
}
