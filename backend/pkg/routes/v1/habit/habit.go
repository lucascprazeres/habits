package habit

import (
	"github.com/gin-gonic/gin"
	"habits/internal/database"
	habitsRepository "habits/internal/database/repositories/habit"
	habitsSchema "habits/internal/schemas/habit"
	habitsService "habits/internal/services/habit"
	"net/http"
)

var service habitsService.Service

func Register(router *gin.RouterGroup) {
	router.POST("", createHabit)

	repository := habitsRepository.NewRepository(database.DB)
	service = habitsService.NewService(repository)
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
