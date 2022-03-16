package accounts

import (
	"net/http"
	"time"

	"example.com/simplebank/database"
	"example.com/simplebank/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var newAccount models.Accounts

	if err := c.BindJSON(&newAccount); err != nil {

		c.Status(http.StatusBadRequest)
	} else {
		newAccount.ID = time.Now().Nanosecond() * time.Now().Second() * time.Now().Minute() * time.Now().Hour() * time.Now().Day() * time.Now().Year()
		if err := database.PostAccounts(&newAccount); err != nil {

			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}

	}

}
