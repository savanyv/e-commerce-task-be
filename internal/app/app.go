package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/config"
	"github.com/savanyv/e-commerce-task-be/internal/database"
	"github.com/savanyv/e-commerce-task-be/internal/delivery/routes"
)

type Server struct {
	app *fiber.App
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		app: fiber.New(),
		config: config,
	}
}

func (s *Server) Run() error {
	_, err := database.InitPostgres(s.config)
	if err != nil {
		log.Println("Failed to connect to database")
		return err
	}

	database.AutoMigrate()

	if err := routes.InitRoutes(s.app); err != nil {
		log.Fatal("Failed to initialize routes")
		return err
	}

	if err := s.app.Listen(":" + s.config.PortServer); err != nil {
		log.Fatal("Failed to start server")
		return err
	}

	log.Println("Server started on port " + s.config.PortServer)
	return nil
}
