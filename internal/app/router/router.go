package router

import (
	"github.com/Anthya1104/gin-base-service/internal/app/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.HealthHandler)

	r.GET("/limiter", handler.RateLimiterHandler)

	return r
}
