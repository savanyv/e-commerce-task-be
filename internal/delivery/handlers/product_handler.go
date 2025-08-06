package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
	v *helpers.Validator
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
		v: helpers.NewValidator(),
	}
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	res, err := h.usecase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Products retrieved successfully",
		"data": res,
	})
}

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := h.usecase.GetProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product retrieved successfully",
		"data": res,
	})
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var req dtos.ProductRequest
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

	res, err := h.usecase.CreateProduct(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"data": res,
	})
}
