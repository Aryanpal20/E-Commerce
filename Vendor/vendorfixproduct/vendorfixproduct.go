package vendoraccess

import (
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"net/http"
	"strings"
)

func ProductVendorAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can split the token and decode the tokens
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	var product pro.Product
	var store pro.Store
	var err user.Error
	// here we can compare with manager
	if role == "vendor" {
		json.NewDecoder(r.Body).Decode(&product)
		database.Database.Where("id = ?", product.Category_Id).Find(&store)
		w.WriteHeader(http.StatusCreated)
		database.Database.Create(&product)
		json.NewEncoder(w).Encode(product)

	} else {
		err = user.Error{Message: "you can't access !!!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
