package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, DB2 AI"})
}

func NotFoundError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Page not found"})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
}
