package main

import (
	"golang-rest-api-mysql/models"
	"golang-rest-api-mysql/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	models.ConnectDatabase()

	route.GET("/api/products", productcontroller.Index)
	route.GET("/api/products/:id", productcontroller.Show)
	route.POST("/api/products", productcontroller.Create)
	route.PUT("/api/products/:id", productcontroller.Update)
	route.DELETE("/api/products/:id", productcontroller.Delete)

	route.Run()

}