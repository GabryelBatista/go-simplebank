package entries

import (
	"net/http"
	"time"

	"example.com/simplebank/database"
	"example.com/simplebank/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var newEntries models.Entries

	if err := c.BindJSON(&newEntries); err != nil {
		c.Status(http.StatusBadRequest)

	} else {
		newEntries.ID = time.Now().Nanosecond() * time.Now().Second() * time.Now().Minute() * time.Now().Hour() * time.Now().Day() * time.Now().Year()
		if err := database.PostEntries(&newEntries); err != nil {
			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)

		}
	}

}
