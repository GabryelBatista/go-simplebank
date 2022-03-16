package entries

import (
	"net/http"
	"strconv"

	"example.com/simplebank/database"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	id := c.Param("id")

	if idAccount, err := strconv.Atoi(id); err != nil {

		c.Status(http.StatusBadRequest)
	} else {
		if result, err := database.GetEntries(idAccount); err != nil {

			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}

}
