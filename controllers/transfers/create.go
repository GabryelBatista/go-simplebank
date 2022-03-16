package transfers

import (
	"net/http"
	"time"

	"example.com/simplebank/database"
	"example.com/simplebank/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var newTransfers models.Transfers

	if err := c.BindJSON(&newTransfers); err != nil {
		c.Status(http.StatusBadRequest)
	}else{
		newTransfers.ID = time.Now().Nanosecond() * time.Now().Second() * time.Now().Minute() * time.Now().Hour() * time.Now().Day() * time.Now().Year()
		if err := database.PostTransfers(&newTransfers); err != nil {
			c.Status(http.StatusInternalServerError)
		}else{
			c.Status(http.StatusOK)
		}
	}

}