package usermodel

import (
	store "gin/models/product_model"
	"time"
)

type User struct {
	ID         int           `json:"id"`
	Email      string        `json:"email"`
	Password   string        `jsson:"password"`
	First_Name string        `json:"first_name"`
	Last_Name  string        `json:"last_name"`
	Role       string        `json:"role"`
	CreatedAT  time.Time     `json:"createdat"`
	Stores     []store.Store `gorm:"ForeignKey:Vendor_Id"`
}
