package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/internal/database"
	"github.com/savanyv/e-commerce-task-be/internal/delivery/handlers"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

func productRoutes(app fiber.Router) {
	repo := repository.NewProductRepository(database.DB)
	variantRepo := repository.NewVariantRepository(database.DB)
	usecase := usecase.NewProductUsecase(repo, variantRepo)
	handler := handlers.NewProductHandler(usecase)

	app.Get("/products", handler.GetAll)
	app.Get("/products/:id", handler.GetByID)
	app.Post("/products", handler.Create)
}
