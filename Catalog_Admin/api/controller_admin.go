package main

import (
	"Catalog_Admin/core"
	"Catalog_Admin/core/interfaces"
	"github.com/gin-gonic/gin"
)
var jsonContentType = "application/json; charset=utf-8"
var catalogRepository interfaces.ICatalogRepository

func RouteOrders(router *gin.Engine)  {
	router.POST("/catalog_admin", CreateProduct)
}


func CreateProduct(c *gin.Context)  {
	product := &core.Product{}
	if catalogRepository.CreateProduct(*product) {
		c.Data(200, jsonContentType, []byte("Created \n"))
	}
	c.Data(500, jsonContentType, []byte("Failed!"))
}


