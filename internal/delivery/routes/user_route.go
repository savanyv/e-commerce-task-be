package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/internal/database"
	"github.com/savanyv/e-commerce-task-be/internal/delivery/handlers"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

func userRoutes(app fiber.Router) {
	repo := repository.NewUserRepository(database.DB)
	usecase := usecase.NewUserUsecase(repo)
	handler := handlers.NewUserHandler(usecase)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}
