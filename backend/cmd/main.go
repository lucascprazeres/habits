package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"habits/internal/database"
	"habits/internal/settings"
	"habits/pkg/routes"
)

func Execute() error {
	r := gin.Default()

	if err := settings.LoadEnvs(); err != nil {
		return err
	}

	if err := database.Connect(); err != nil {
		return err
	}

	routes.Register(r)

	envs := settings.GetEnvs().App

	addr := fmt.Sprintf(":%s", envs.Port)
	if err := r.Run(addr); err != nil {
		return err
	}

	return nil
}
