package routes

import (
	"example.com/simplebank/controllers/accounts"
	"github.com/gin-gonic/gin"
)


func Accounts(router *gin.Engine)  {

	endPoint := "/accounts"
	router.GET(endPoint, accounts.GetAll)
	router.GET(endPoint+"/:id", accounts.GetID)
	router.POST(endPoint, accounts.Create)
	router.PUT(endPoint, accounts.Edit)
	router.DELETE(endPoint+"/:id", accounts.Delete)

}