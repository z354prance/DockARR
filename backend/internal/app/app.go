package app

import (
	"github.com/gin-gonic/gin"

	"github.com/z354prance/DockARR/backend/internal/router"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	return &App{
		router: router.New(),
	}
}

func (a *App) Run() error {
	return a.router.Run(":8080")
}
