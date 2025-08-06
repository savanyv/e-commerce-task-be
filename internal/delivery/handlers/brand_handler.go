package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

type BrandHandler struct {
	usecase usecase.BrandUsecase
	v *helpers.Validator
}

func NewBrandHandler(usecase usecase.BrandUsecase) *BrandHandler {
	return &BrandHandler{
		usecase: usecase,
	}
}

func (h *BrandHandler) GetAllBrands(c *fiber.Ctx) error {
	res, err := h.usecase.GetAllBrands()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brands retrieved successfully",
		"data": res,
	})
}

func (h *BrandHandler) CreateBrand(c *fiber.Ctx) error {
	var req dtos.BrandRequest
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

	res, err := h.usecase.CreateBrand(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Brand created successfully",
		"data": res,
	})
}
