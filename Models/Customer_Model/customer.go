package customermodel

// import cart "e-Commerce/Models/Cart_Model"

type Customer struct {
	ID                      int    `json:"id"`
	Product_Id              int    `json:"product_id"`
	Customer_First_Name     string `json:"customer_first_name"`
	Customer_Last_Name      string `json:"customer_last_name"`
	Customer_Address        string `json:"customer_address"`
	Customer_Phone_No       string `gorm:"unique:customer_phone_no"`
	Select_Product_Name     string `json:"select_product_name"`
	Select_Product_Quantity int    `json:"select_product_quantity"`
	Quantity_wise_Rate      int    `json:"quantity_wise_rate"`
	Stock                   string `json:"stock"`
	Customer_Id             int    `json:"customer_id"`
	// Carts                   []cart.Cart `gorm:"ForeignKey:Userid"`
}
