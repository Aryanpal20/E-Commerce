package main

import (
	admin "gin/controller/admin"
	auth "gin/controller/auth"
	cust "gin/controller/customer"
	vendor "gin/controller/vendor"
	"gin/database"
	"gin/middelware"

	"github.com/gin-gonic/gin"
)

func main() {

	database.DataMigration()

	r := gin.Default()

	r.Static("/static", "./static")

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/fixstore", middelware.AuthRequired(), admin.PostAdminStore)
	r.PATCH("/updatestore/:id", middelware.AuthRequired(), admin.UpdateAdminStore)
	r.DELETE("/deletestore/:id", middelware.AuthRequired(), admin.DeleteAdminStore)
	r.POST("/fixcategory", middelware.AuthRequired(), admin.PostAdminCategory)
	r.PATCH("/updatecategory/:id", middelware.AuthRequired(), admin.UpdateAdminCategory)
	r.DELETE("/deletecategory/:id", middelware.AuthRequired(), admin.DeleteAdminCategory)
	r.POST("/fixproduct", middelware.AuthRequired(), vendor.PostVendorProduct)
	r.PATCH("/updateproduct/:id", middelware.AuthRequired(), vendor.UpdateVendorProduct)
	r.DELETE("/deleteproduct/:id", middelware.AuthRequired(), vendor.DeleteVendorProduct)
	r.POST("/customerselectproduct", middelware.AuthRequired(), cust.CustomerSelectProduct)
	r.PUT("/customerupdateselectproduct/:id", middelware.AuthRequired(), cust.CustomerUpdateSelectProduct)
	r.DELETE("/customerdeleteselectproduct/:id", middelware.AuthRequired(), cust.CustomerDeleteSelectProduct)
	r.POST("/customerorder", middelware.AuthRequired(), cust.CustomerOrder)

	r.Run()
}
