package cartcreatewithlogin

import (
	email "e-Commerce/FetchOnlyOneData/fetchemail"
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	cart "e-Commerce/Models/Cart_Model"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token := strings.Split(r.Header["Token"][0], " ")[1]
		fmt.Println(r.Header)
		fmt.Println(r.Header["Cookie"])
		role := role.Is_manager(token)
		email := email.Task_creator(token)
		if role == "customer" {
			var cart cart.Cart
			var user user.User
			var pro pro.Product
			json.NewDecoder(r.Body).Decode(&cart)
			database.Database.Where("id = ?", cart.ProductId).Find(&pro)
			fmt.Println(pro)
			database.Database.Where("id = ?", cart.Userid).Find(&user)
			fmt.Println(user)
			if user.Email == email {
				cart.Select_Product_Name = pro.Product_Name
				cart.Quantity_wise_Rate = pro.Product_Rate
				cart.Userid = user.ID
				if cart.Select_Product_Quantity <= pro.Product_Quantity {
					if cart.Select_Product_Quantity > 1 {
						rate := pro.Product_Rate * cart.Select_Product_Quantity
						cart.Quantity_wise_Rate = rate
					}
				} else {
					cart.Select_Product_Quantity = pro.Product_Quantity
				}
				database.Database.Create(&cart)
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(cart)
			} else {
				err := "you can enter your id"
				json.NewEncoder(w).Encode(err)
			}

		}
	} else {
		err := "First you can login"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)

	}

}
