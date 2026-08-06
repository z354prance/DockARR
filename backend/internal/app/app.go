package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/z354prance/DockARR/backend/internal/config"
	"github.com/z354prance/DockARR/backend/internal/logging"
	"github.com/z354prance/DockARR/backend/internal/router"
)

type App struct {
	router *gin.Engine
	logger *zap.Logger
	config *config.Config
}

func New() (*App, error) {
	logger, err := logging.New()
	if err != nil {
		return nil, err
	}

	cfg := config.Load()

	return &App{
		router: router.New(),
		logger: logger,
		config: cfg,
	}, nil
}

func (a *App) Run() error {
	address := fmt.Sprintf("%s:%s", a.config.Host, a.config.Port)

	a.logger.Info(
		"Starting DockARR",
		zap.String("address", address),
	)

	return a.router.Run(address)
}
