package admin

import (
	"gin/database"
	store "gin/models/product_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAdminStore(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var input store.Store
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := store.Store{Store_Name: input.Store_Name, Store_Address: input.Store_Address, Vendor_Id: input.Vendor_Id}
		database.Database.Create(&store)
		c.JSON(http.StatusAccepted, gin.H{"Data": store})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't access"})
	}

}

func UpdateAdminStore(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var stores store.Store
		var input store.Store
		if err := database.Database.Where("id = ?", c.Param("id")).First(&stores).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := store.Store{Store_Name: input.Store_Name, Store_Address: input.Store_Address}
		// database.Database.Model(&stores).Updates(store)
		c.JSON(http.StatusCreated, gin.H{"Data": store})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't access"})
	}

}

func DeleteAdminStore(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var stores store.Store
		if err := database.Database.Where("id = ?", c.Param("id")).First(&stores).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		// database.Database.Delete(&stores)
		c.JSON(http.StatusContinue, gin.H{"Message": "your store data deleted successfully !!!"})
	}
}

func PostAdminCategory(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var input store.Category
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		category := store.Category{Category_Name: input.Category_Name, Store_Id: input.Store_Id}
		database.Database.Create(&category)
		c.JSON(http.StatusAccepted, gin.H{"Data": category})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can't access"})
	}

}

func UpdateAdminCategory(c *gin.Context) {
	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var categories store.Category
		var input store.Category
		if err := database.Database.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		category := store.Category{Category_Name: input.Category_Name}
		// database.Database.Model(&categories).Updates(category)
		c.JSON(http.StatusCreated, gin.H{"Data": category})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't access"})
	}

}

func DeleteAdminCategory(c *gin.Context) {

	tokenrole := c.GetString("role")
	if tokenrole == "admin" {
		var categories store.Category
		if err := database.Database.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		// database.Database.Delete(&categories)
		c.JSON(http.StatusContinue, gin.H{"Message": "your category data deleted successfully !!!"})
	}
}
