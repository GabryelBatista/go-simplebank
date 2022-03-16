package accounts

import (
	"net/http"
	"strconv"

	"example.com/simplebank/database"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	
	id := c.Param("id")

	if idAccount, err := strconv.Atoi(id); err != nil {

		c.Status(http.StatusBadRequest)
	} else {
		if err := database.DeleteAccounts(idAccount); err != nil {

			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
	}
}
