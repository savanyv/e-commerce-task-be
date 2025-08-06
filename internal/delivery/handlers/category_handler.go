package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

type CategoryHandler struct {
	usecsae usecase.CategoryUsecase
	v *helpers.Validator
}

func NewCategoryHandler(usecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecsae: usecase,
		v: helpers.NewValidator(),
	}
}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	res, err := h.usecsae.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Categories retrieved successfully",
		"data": res,
	})
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req dtos.CategoryRequest
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

	res, err := h.usecsae.CreateCategory(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
		"data": res,
	})
}
