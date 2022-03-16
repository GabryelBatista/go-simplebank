package routes

import (
	"example.com/simplebank/controllers/entries"
	"github.com/gin-gonic/gin"
)

func Entries(router *gin.Engine)  {

	endPoint := "/entries"
	router.GET(endPoint+"/:id", entries.GetAll)
	router.POST(endPoint, entries.Create)
}