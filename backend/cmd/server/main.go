package main

import (
	"log"

	"github.com/z354prance/DockARR/backend/internal/app"
)

func main() {
	log.Println("Starting DockARR API on :8080")

	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}
