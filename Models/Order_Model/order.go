package ordermodel

type Order struct {
	ID               int    `json:"id"`
	Product_Id       int    `json:"product_id"`
	Product_Name     string `json:"product_name"`
	Product_Quantity int    `json:"product_quantity"`
	Total_Rate       int    `json:"total_rate"`
	Status           string `json:"status"`
	OrderID          int    `json:"orderid"`
	// Userid           int    `json:"userid"`
}
