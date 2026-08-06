package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/z354prance/DockARR/backend/internal/logging"
	"github.com/z354prance/DockARR/backend/internal/router"
)

type App struct {
	router *gin.Engine
	logger *zap.Logger
}

func New() (*App, error) {

	logger, err := logging.New()
	if err != nil {
		return nil, err
	}

	return &App{
		router: router.New(),
		logger: logger,
	}, nil
}

func (a *App) Run() error {

	a.logger.Info("DockARR starting")

	return a.router.Run(":8080")
}
