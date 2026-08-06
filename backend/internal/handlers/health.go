package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/z354prance/DockARR/backend/internal/version"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"application": version.AppName,
		"status":      "ok",
		"version":     version.Version,
	})
}
