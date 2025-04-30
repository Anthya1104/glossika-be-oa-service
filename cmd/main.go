package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yourname/base-service/internal/handler"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    r.GET("/limiter", handler.RateLimiterHandler)

    r.Run(":8080")
}
