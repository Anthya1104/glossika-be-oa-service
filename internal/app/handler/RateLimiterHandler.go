package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RateLimiterHandler(c *gin.Context) {
	// RateLimiter Simulating
	ip := c.ClientIP()
	c.JSON(http.StatusOK, gin.H{
		"message": "Request allowed",
		"ip":      ip,
	})
}
