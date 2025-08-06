package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/internal/database"
	"github.com/savanyv/e-commerce-task-be/internal/delivery/handlers"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

func categoryRoutes(app fiber.Router) {
	repo := repository.NewCategoryRepository(database.DB)
	usecase := usecase.NewCategoryUsecase(repo)
	handler := handlers.NewCategoryHandler(usecase)

	app.Get("/categories", handler.GetAllCategories)
	app.Post("/categories", handler.CreateCategory)
}
