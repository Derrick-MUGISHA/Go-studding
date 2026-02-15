package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []gin.H{
		{"id": 1, "name": "Derrick"},
		{"id": 2, "name": "Alice"},
	}

	c.JSON(http.StatusOK, users)
}
