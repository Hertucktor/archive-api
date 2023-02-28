package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
}
