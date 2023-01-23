package customerselect

import (
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	cust "e-Commerce/Models/Customer_Model"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Customer_Select(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	if role == "customer" {
		var cust cust.Customer
		var user user.User
		var pro pro.Product
		json.NewDecoder(r.Body).Decode(&cust)
		database.Database.Where("id = ?", cust.Product_Id).Find(&pro)
		fmt.Println(pro)
		database.Database.Where("id = ?", cust.Customer_Id).Find(&user)
		fmt.Println(user)
		cust.Customer_First_Name = user.First_Name
		cust.Customer_Last_Name = user.Last_Name
		cust.Customer_Phone_No = user.Phone
		cust.Customer_Address = user.Address
		cust.Select_Product_Name = pro.Product_Name
		cust.Quantity_wise_Rate = pro.Product_Rate
		cust.Stock = strconv.Itoa(pro.Product_Quantity) + " Products in stock"
		if cust.Select_Product_Quantity <= pro.Product_Quantity {
			if cust.Select_Product_Quantity > 1 {
				rate := pro.Product_Rate * cust.Select_Product_Quantity
				cust.Quantity_wise_Rate = rate
			}
		} else {
			avilable := "Only " + strconv.Itoa(pro.Product_Quantity) + " in stock"
			cust.Stock = avilable
			cust.Select_Product_Quantity = pro.Product_Quantity
			// json.NewEncoder(w).Encode(cust)
		}
		database.Database.Create(&cust)
		json.NewEncoder(w).Encode(&cust)

		// database.Database.Create(&cust)
		// json.NewEncoder(w).Encode(&cust)
		// } else {
		// 	if err := database.Database.Model(&cust).Where("id = ?", cust.Product_Id).Update("select_product_quantity", cust.Select_Product_Quantity).Error; err != nil {
		// 		fmt.Printf("update err != nil; %v\n", err)
		// 	}
		// 	database.Database.Save(&custs)
		// 	json.NewEncoder(w).Encode(custs)

	} else {
		role := "You Can't Access"
		json.NewEncoder(w).Encode(role)
	}
}
