package main

import (
	"Catalog_User/core/interfaces"
	"github.com/gin-gonic/gin"
)
var jsonContentType = "application/json; charset=utf-8"
var catalogRepository interfaces.ICatalogRepository

func RouteOrders(router *gin.Engine)  {
	router.GET("/catalog_user", GetAll)

}

func GetAll(c *gin.Context)  {
	catalog := catalogRepository.GetAll()
	c.JSON(200, catalog)
}


