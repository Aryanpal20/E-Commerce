package usermodel

import (
	customer "e-Commerce/Models/Customer_Model"
	store "e-Commerce/Models/Product_Model"
)

type User struct {
	ID         int                 `json:"id"`
	Email      string              `json:"email"`
	Password   string              `json:"password"`
	First_Name string              `json:"first_name"`
	Last_Name  string              `json:"last_name"`
	Phone      string              `gorm:"unique:phone"`
	Address    string              `json:"address"`
	Role       string              `json:"role"`
	Stores     []store.Store       `gorm:"ForeignKey:Vendor_Id"`
	Customers  []customer.Customer `gorm:"ForeignKey:Customer_Id"`
}
