package customer

import (
	"gin/database"
	cust "gin/models/customer_model"
	order "gin/models/order_model"
	store "gin/models/product_model"
	entity "gin/models/user_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerSelectProduct(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "customer" {
		var user entity.User
		tokenid := c.GetFloat64("id")
		tokenemail := c.GetString("email")
		database.Database.Where("id = ?", tokenid).Find(&user)
		if tokenemail == user.Email {
			var input cust.Customer
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var product store.Product
			pro := cust.Customer{Product_Id: input.Product_Id}
			database.Database.Where("id = ?", pro.Product_Id).Find(&product)
			products := cust.Customer{Product_Id: input.Product_Id, Customer_First_Name: user.First_Name, Customer_Last_Name: user.Last_Name,
				Select_Product_Name: product.Product_Name, Select_Product_Quantity: input.Select_Product_Quantity,
				Quantity_wise_Rate: product.Product_Rate}
			if products.Select_Product_Quantity <= product.Product_Quantity {
				if products.Select_Product_Quantity > 1 {
					rate := product.Product_Rate * products.Select_Product_Quantity
					products.Quantity_wise_Rate = rate
				}
			}
			// database.Database.Create(&products)
			c.JSON(http.StatusCreated, gin.H{"Data": products})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please login your account"})
	}
}

func CustomerUpdateSelectProduct(c *gin.Context) {
	tokenrole := c.GetString("role")
	if tokenrole == "customer" {
		var user entity.User
		tokenid := c.GetFloat64("id")
		tokenemail := c.GetString("email")
		database.Database.Where("id = ?", tokenid).Find(&user)
		if tokenemail == user.Email {
			var input cust.Customer
			var customer cust.Customer
			if err := database.Database.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
				return
			}
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var product store.Product
			pro := cust.Customer{Product_Id: input.Product_Id}
			database.Database.Where("id = ?", pro.Product_Id).Find(&product)
			products := cust.Customer{Product_Id: input.Product_Id, Customer_First_Name: user.First_Name, Customer_Last_Name: user.Last_Name,
				Select_Product_Name: product.Product_Name, Select_Product_Quantity: input.Select_Product_Quantity,
				Quantity_wise_Rate: product.Product_Rate}
			if products.Select_Product_Quantity <= product.Product_Quantity {
				if products.Select_Product_Quantity > 1 {
					rate := product.Product_Rate * products.Select_Product_Quantity
					products.Quantity_wise_Rate = rate
				}
			}
			database.Database.Model(&customer).Updates(products)
			c.JSON(http.StatusAccepted, gin.H{"Data": products})
		}
	}
}

func CustomerDeleteSelectProduct(c *gin.Context) {
	tokenrole := c.GetString("role")
	if tokenrole == "customer" {
		var user entity.User
		tokenid := c.GetFloat64("id")
		tokenemail := c.GetString("email")
		database.Database.Where("id = ?", tokenid).Find(&user)
		if tokenemail == user.Email {
			var customer cust.Customer
			if err := database.Database.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
				return
			}
			// database.Database.Delete(&customer)
			c.JSON(http.StatusAccepted, gin.H{"Message": "your selected product data deleted successfully !!!"})
		}
	}
}

func CustomerOrder(c *gin.Context) {
	tokenrole := c.GetString("role")
	if tokenrole == "customer" {
		var user entity.User
		tokenid := c.GetFloat64("id")
		tokenemail := c.GetString("email")
		database.Database.Where("id = ?", tokenid).Find(&user)
		if tokenemail == user.Email {
			var input order.Order
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var customer cust.Customer
			orders := order.Order{CustomerID: input.CustomerID}
			database.Database.Where("id = ?", orders.CustomerID).Find(&customer)
			order := order.Order{CustomerID: orders.CustomerID, Product_Id: customer.Product_Id, Product_Name: customer.Select_Product_Name,
				Product_Quantity: customer.Select_Product_Quantity, Total_Rate: customer.Quantity_wise_Rate, Status: input.Status}
			database.Database.Create(&order)
			c.JSON(http.StatusAccepted, gin.H{"Data": order})
		}
	}
}
