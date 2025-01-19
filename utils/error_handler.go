package utils

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Response is a common response struct

func ErrorHandlerMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			statusCode := fiber.StatusInternalServerError
			var message string

			switch e := err.(type) {
			case *fiber.Error:
				statusCode = e.Code
				message = e.Message
			default:
				message = err.Error()
			}

			logger.Error("Error occurred", zap.Error(err))

			c.Response().SetStatusCode(statusCode)

			response := ResponseError{
				Code:    statusCode,
				Status:  "error",
				Message: message,
			}

			return c.Status(statusCode).JSON(response)
		}

		return nil
	}
}
