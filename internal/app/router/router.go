package router

import (
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/handler"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", handler.HealthHandler)

	r.GET("api/users", handler.GetUserInfoAPI)

	r.POST("api/v1/users", handler.UserRegisterHandler)

	r.GET("/api/v1/users/verify", handler.UserActivateHandler)

	r.POST("/api/v1/auth/login", handler.UserLoginHandler)

	// auth API group
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWTAuthMiddleware())
	// TODO: remove this route after testing
	// auth.GET("/test-api", handler.TryTestHandler)

	return r
}

func Setup() error {
	Router = SetupRouter()

	return nil
}
