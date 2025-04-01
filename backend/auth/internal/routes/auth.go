package routes

import "github.com/gin-gonic/gin"

func SetupAuthRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from auth"})
	})
}
