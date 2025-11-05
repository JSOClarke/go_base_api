package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is up and running :)"})
}
