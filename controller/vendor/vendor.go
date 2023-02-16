package vendor

import (
	"gin/database"
	pro "gin/models/product_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostVendorProduct(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "vendor" {
		var input pro.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := pro.Product{Product_Name: input.Product_Name, Product_Rate: input.Product_Rate, Product_Quantity: input.Product_Quantity, Category_Id: input.Category_Id}
		database.Database.Create(&product)
		c.JSON(http.StatusAccepted, gin.H{"Data": product})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't access"})
	}
}

func UpdateVendorProduct(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "vendor" {
		var products pro.Product
		var input pro.Product
		if err := database.Database.Where("id = ?", c.Param("id")).First(&products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := pro.Product{Product_Rate: input.Product_Rate, Product_Quantity: input.Product_Quantity}
		// database.Database.Model(&products).Updates(product)
		c.JSON(http.StatusCreated, gin.H{"Data": product})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't access"})
	}

}

func DeleteVendorProduct(c *gin.Context) {
	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var products pro.Product
		if err := database.Database.Where("id = ?", c.Param("id")).First(&products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		// database.Database.Delete(&products)
		c.JSON(http.StatusContinue, gin.H{"Message": "your product data deleted successfully !!!"})
	}
}
