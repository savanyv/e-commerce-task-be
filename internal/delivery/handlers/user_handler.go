package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
	"github.com/savanyv/e-commerce-task-be/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
	v *helpers.Validator
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
		v: helpers.NewValidator(),
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dtos.RegisterRequest
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

	user, err := h.usecase.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data": user,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dtos.LoginRequest
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

	user, err := h.usecase.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
		"data": user,
	})
}
