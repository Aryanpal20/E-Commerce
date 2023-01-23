package database

import (
	cart "e-Commerce/Models/Cart_Model"
	cust "e-Commerce/Models/Customer_Model"
	pro "e-Commerce/Models/Product_Model"
	user "e-Commerce/Models/User_Model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"

var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())

		panic("connection failed")
	}
	Database.AutoMigrate(user.User{}, pro.Product{}, pro.Store{}, pro.Category{}, cust.Customer{}, cart.Cart{})
}
