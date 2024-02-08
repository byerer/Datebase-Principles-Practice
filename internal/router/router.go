package router

import (
	"GradingSystem/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", register)
	r.POST("/forgetPassword", forgetPassword)
	r.GET("/login", login)
	r.GET("/sendCode", sendCode)

	api := r.Group("/api", middleware.JWTAuthentication())
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "test",
			})
		})
	}
	_ = r.Run()
}
