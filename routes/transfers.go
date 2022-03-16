package routes

import (
	"example.com/simplebank/controllers/transfers"
	"github.com/gin-gonic/gin"
)


func Transfers( router *gin.Engine)  {
	endPoint := "/transfers"
	router.GET(endPoint+"/:id", transfers.GetAll)
	router.POST(endPoint, transfers.Create)
}