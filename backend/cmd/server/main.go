package main

import (
	"log"

	"github.com/z354prance/DockARR/backend/internal/router"
)

func main() {
	r := router.New()

	log.Println("Starting DockARR API on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
