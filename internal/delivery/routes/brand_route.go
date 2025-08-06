package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/internal/database"
	"github.com/savanyv/e-commerce-task-be/internal/delivery/handlers"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

func brandRoutes(app fiber.Router) {
	repo := repository.NewBrandRepository(database.DB)
	usecase := usecase.NewBrandUsecase(repo)
	handler := handlers.NewBrandHandler(usecase)

	app.Get("/brands", handler.GetAllBrands)
	app.Post("/brands", handler.CreateBrand)
}
