package adminchangescategoryvalue

import (
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CategoryUpdateByAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can split the token and decode the tokens
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	var err user.Error
	// here we can compare with manager
	if role == "admin" {
		var category pro.Category
		// here we can find the from database by id
		database.Database.First(&category, mux.Vars(r)["store_id"])
		json.NewDecoder(r.Body).Decode(&category)
		if err := database.Database.Model(&category).Where("store_id = ?", category.Store_Id).Update("category_name", category.Category_Name).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		w.WriteHeader(http.StatusOK)
		database.Database.Save(&category)
		json.NewEncoder(w).Encode(category)

	} else {
		err = user.Error{Message: "you can't access !!!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
