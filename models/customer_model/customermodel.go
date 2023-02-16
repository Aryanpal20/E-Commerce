package customermodel

import order "gin/models/order_model"

type Customer struct {
	ID                      int           `json:"id"`
	Product_Id              int           `json:"product_id"`
	Customer_First_Name     string        `json:"customer_first_name"`
	Customer_Last_Name      string        `json:"customer_last_name"`
	Select_Product_Name     string        `json:"select_product_name"`
	Select_Product_Quantity int           `json:"select_product_quantity"`
	Quantity_wise_Rate      int           `json:"quantity_wise_rate"`
	Orders                  []order.Order `gorm:"ForeignKey:CustomerID"`
}
