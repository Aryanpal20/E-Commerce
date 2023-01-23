package cartmodel

type Cart struct {
	ID                      int    `json:"id"`
	Select_Product_Name     string `json:"select_product_name"`
	Select_Product_Quantity int    `json:"select_product_quantity"`
	Quantity_wise_Rate      int    `json:"quantity_wise_rate"`
	ProductId               int    `json:"productid"`
	Userid                  int    `json:"userid"`
}
