package middleware

import (
	e "notification-service/internal/exception"

	"github.com/gofiber/fiber/v2"
)

func ErrorControllerMiddleware(c *fiber.Ctx) error {

	err := c.Next()

	if err != nil {
		return e.HandleHttpErrorFiber(c, err)
	}

	return nil
}
