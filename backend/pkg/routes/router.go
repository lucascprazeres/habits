package routes

import (
	"github.com/gin-gonic/gin"
	habitsV1 "habits/pkg/routes/v1/habit"
)

func Register(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/health-check", healthCheck)

	habitsV1.Register(v1.Group("/habits"))
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{})
}
