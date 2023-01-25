package adminchangestorevalue

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

func StoreUpdateByAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can split the token and decode the tokens
	token := strings.Split(r.Header["Token"][0], " ")[1]
	role := role.Is_manager(token)
	var err user.Error
	// here we can compare with manager
	if role == "admin" {
		var store pro.Store
		// here we can find the from database by id
		database.Database.First(&store, mux.Vars(r)["vendor_id"])
		json.NewDecoder(r.Body).Decode(&store)
		if err := database.Database.Model(&store).Where("id = ?", store.Vendor_Id).Update("store_name", store.Store_Name).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		if err := database.Database.Model(&store).Where("id = ?", store.Vendor_Id).Update("store_address", store.Store_Address).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		w.WriteHeader(http.StatusOK)
		database.Database.Save(&store)
		json.NewEncoder(w).Encode(store)

	} else {
		err = user.Error{Message: "you can't access !!!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
