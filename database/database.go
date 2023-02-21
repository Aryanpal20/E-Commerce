package database

import (
	"fmt"
	cust "gin/models/customer_model"
	order "gin/models/order_model"
	product "gin/models/product_model"
	user "gin/models/user_model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

// var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/e_commerce?parseTime=true"

var err error

func DataMigration() {
	// DB, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	DB, err := gorm.Open(sqlite.Open("DB"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	DB.AutoMigrate(user.User{}, product.Category{}, product.Product{}, product.Store{}, cust.Customer{}, order.Order{})
	Database = DB
}
