package customerchange

import (
	email "e-Commerce/FetchOnlyOneData/fetchemail"
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	cust "e-Commerce/Models/Customer_Model"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CustomerChange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	email := email.Task_creator(token)
	role := role.Is_manager(token)
	if role == "customer" {
		var user user.User
		var cust cust.Customer
		var pro pro.Product
		database.Database.First(&user, mux.Vars(r)["customer_id"])
		fmt.Println(user)
		database.Database.Where("customer_id = ?", user.ID).Find(&cust)
		fmt.Println(cust)
		database.Database.Where("id = ?", cust.Product_Id).Find(&pro)
		// fmt.Println(user)
		// fmt.Println(user.Email)
		// fmt.Println(email)
		if email == user.Email {

			json.NewDecoder(r.Body).Decode(&cust)
			if err := database.Database.Model(&cust).Where("customer_id = ?", cust.Customer_Id).Update("select_product_quantity", cust.Select_Product_Quantity).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			rate := pro.Product_Rate * cust.Select_Product_Quantity
			cust.Quantity_wise_Rate = rate
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
