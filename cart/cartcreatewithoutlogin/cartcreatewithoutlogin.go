package cartcreatewithoutlogin

import (
	cart "e-Commerce/Models/Cart_Model"
	pro "e-Commerce/Models/Product_Model"
	"e-Commerce/database"
	"encoding/json"
	"net/http"
)

func AddToCartWithoutLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart cart.Cart
	var pro pro.Product
	json.NewDecoder(r.Body).Decode(&cart)
	database.Database.Where("id = ?", cart.ProductId).Find(&pro)
	cart.Select_Product_Name = pro.Product_Name
	cart.Quantity_wise_Rate = pro.Product_Rate
	if cart.Select_Product_Quantity <= pro.Product_Quantity {
		if cart.Select_Product_Quantity > 1 {
			rate := pro.Product_Rate * cart.Select_Product_Quantity
			cart.Quantity_wise_Rate = rate
		}
	} else {
		cart.Select_Product_Quantity = pro.Product_Quantity
	}
	database.Database.Create(&cart)
	json.NewEncoder(w).Encode(cart)

}
