package vendorchangesproductvalue

import (
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	pro "e-Commerce/Models/Product_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func ProductUpdateByVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can split the token and decode the tokens
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)

	// here we can compare with manager
	if role == "vendor" {
		var product pro.Product
		// here we can find the from database by id
		database.Database.First(&product, mux.Vars(r)["category_id"])
		json.NewDecoder(r.Body).Decode(&product)
		if err := database.Database.Model(&product).Where("category_id = ?", product.Category_Id).Update("product_name", product.Product_Name).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		if err := database.Database.Model(&product).Where("category_id = ?", product.Category_Id).Update("product_rate", product.Product_Rate).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		if err := database.Database.Model(&product).Where("category_id = ?", product.Category_Id).Update("product_quantity", product.Product_Quantity).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		w.WriteHeader(http.StatusOK)
		database.Database.Save(&product)
		json.NewEncoder(w).Encode(product)

	} else {
		err := "you can't access !!!"
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
