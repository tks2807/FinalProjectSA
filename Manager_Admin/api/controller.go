package api

import (
	"Manager_Admin/core"
	"Manager_Admin/core/interfaces"
	"github.com/gin-gonic/gin"
)

var jsonContentType = "application/json; charset=utf-8"
var accountRepository interfaces.IAccountRepository

func RouteOrders(router *gin.Engine)  {
	router.POST("/account", CreateAccount)
	router.GET("/accounts", GetAll)
}


func CreateAccount(c *gin.Context)  {
	acc := &core.Account{}
	if accountRepository.CreateAccount(*acc) {
		c.Data(200, jsonContentType, []byte("Created!\n"))
	}
	c.Data(500, jsonContentType, []byte("ERROR!"))
}

func GetAll(c *gin.Context)  {
	acc := accountRepository.GetAll()
	c.JSON(200, acc)
}
