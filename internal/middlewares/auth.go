package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
)

func AuthMiddlewares(jwtService helpers.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		token := bearerToken[1]
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}
