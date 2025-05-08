package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
