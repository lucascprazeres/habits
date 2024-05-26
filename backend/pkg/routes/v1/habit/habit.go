package habit

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	habitsSchema "habits/internal/schemas/habit"
	habitsService "habits/internal/services/habit"
	"net/http"
)

var service habitsService.Service

func Register(router *gin.RouterGroup) {
	service = habitsService.NewService()

	router.POST("", createHabit)
	router.PATCH(":id/toggle", toggleHabit)
}

func createHabit(c *gin.Context) {
	var input habitsSchema.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := service.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func toggleHabit(c *gin.Context) {
	i := c.Param("id")
	id, err := uuid.Parse(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit, err := service.Toggle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habit)
}
