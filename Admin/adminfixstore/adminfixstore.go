package adminfix

import (
	role "e-Commerce/FetchOnlyOneData/fetchrole"
	pro "e-Commerce/Models/Product_Model"
	entity "e-Commerce/Models/User_Model"
	"e-Commerce/database"
	"encoding/json"
	"net/http"
	"strings"
)

func StoreAminAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can split the token and decode the tokens
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	// email := tokenEmail
	// var product_name string
	var store pro.Store
	var user entity.User
	// here we can compare with manager
	if role == "admin" {
		json.NewDecoder(r.Body).Decode(&store)
		database.Database.Where("id = ?", store.Vendor_Id).Find(&user)
		w.WriteHeader(http.StatusCreated)
		database.Database.Create(&store)
		json.NewEncoder(w).Encode(store)

	} else {
		err := "you can't access !!!"
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
