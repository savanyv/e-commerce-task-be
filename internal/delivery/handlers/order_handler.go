package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
	v *helpers.Validator
}

func NewOrderHandler(usecase usecase.OrderUsecase) OrderHandler {
	return OrderHandler{
		usecase: usecase,
		v: helpers.NewValidator(),
	}
}

func (h *OrderHandler) Checkout(c *fiber.Ctx) error {
	var req dtos.OrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.v.Validate(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := h.usecase.Checkout(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
		"data": res,
	})
}
