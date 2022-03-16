package accounts

import (
	"net/http"

	"example.com/simplebank/database"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	if result, err := database.GetAccounts(); err != nil {

		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
