package main

import (
	"codebase/jwt/handler"
	"codebase/jwt/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(middleware.Logger())
	r.POST("/login", handler.Login)

	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	authGroup.GET("/user_info", handler.GetUserInfo)
	r.Run("localhost:7001")
}
