package adminfixcategory

import (
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CategoryAminAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can split the token and decode the tokens
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	var err user.Error
	var category pro.Category
	var product pro.Product
	// here we can compare with admin
	if role == "admin" {
		json.NewDecoder(r.Body).Decode(&category)
		database.Database.Where("id = ?", category.Store_Id).Find(&product)
		fmt.Println(product)

		w.WriteHeader(http.StatusCreated)
		database.Database.Create(&category)
		json.NewEncoder(w).Encode(category)

	} else {
		err = user.Error{Message: "you can't access !!!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
