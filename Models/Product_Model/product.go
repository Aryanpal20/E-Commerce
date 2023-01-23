package productmodel

import (
	cart "e-Commerce/Models/Cart_Model"
	customer "e-Commerce/Models/Customer_Model"
)

type Store struct {
	ID            int        `json:"id"`
	Store_Name    string     `json:"store_name"`
	Store_Address string     `json:"store_address"`
	Vendor_Id     int        `json:"vendor_id"`
	Categorys     []Category `gorm:"ForeignKey:Store_Id"`
}

type Category struct {
	ID            int       `json:"id"`
	Category_Name string    `json:"category_name"`
	Store_Id      int       `json:"store_id"`
	Products      []Product `gorm:"ForeignKey:Category_Id"`
}

type Product struct {
	ID               int                 `json:"id"`
	Product_Name     string              `json:"product_name"`
	Product_Rate     int                 `json:"product_rate"`
	Product_Quantity int                 `json:"product_quantity"`
	Category_Id      int                 `json:"category_id"`
	Customers        []customer.Customer `gorm:"ForeignKey:Product_Id"`
	Carts            []cart.Cart         `gorm:"ForeignKey:ProductId"`
}
