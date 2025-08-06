package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/internal/database"
	"github.com/savanyv/e-commerce-task-be/internal/delivery/handlers"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

func orderRoutes(app fiber.Router) {
	repo := repository.NewOrderRepository(database.DB)
	variantRepo := repository.NewVariantRepository(database.DB)
	usecase := usecase.NewOrderUsecase(repo, variantRepo)
	handler := handlers.NewOrderHandler(usecase)

	app.Post("/checkout", handler.Checkout)
}
