package main

import (
	"log"

	"github.com/savanyv/e-commerce-task-be/config"
	"github.com/savanyv/e-commerce-task-be/internal/app"
)

func main() {
	config := config.LoadConfig()

	app := app.NewServer(config)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
