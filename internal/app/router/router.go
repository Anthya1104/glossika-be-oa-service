package router

import (
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/handler"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.HealthHandler)

	r.GET("api/users", handler.GetUserInfoAPI)

	r.GET("/limiter", handler.RateLimiterHandler)

	return r
}

func Setup() error {
	Router = SetupRouter()

	return nil
}
