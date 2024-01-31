package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}