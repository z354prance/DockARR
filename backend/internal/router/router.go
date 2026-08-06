package router

import (
	"github.com/gin-gonic/gin"

	"github.com/z354prance/DockARR/backend/internal/handlers"
)

func New() *gin.Engine {
	r := gin.New()

	api := r.Group("/api/v1")
	{
		api.GET("/health", handlers.Health)
	}

	return r
}
