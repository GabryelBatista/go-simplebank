package accounts

import (
	"net/http"

	"example.com/simplebank/database"
	"example.com/simplebank/models"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	var newAccount models.Accounts

	if err := c.BindJSON(&newAccount); err != nil {

		c.Status(http.StatusBadRequest)
	} else {
		if err := database.PutAccounts(&newAccount); err != nil {

			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
	}
}
