package main

import (
	"example.com/simplebank/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Accounts(router)
	routes.Entries(router)
	routes.Transfers(router)
	
	router.Run()
}