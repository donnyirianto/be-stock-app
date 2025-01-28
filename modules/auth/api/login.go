package api

import (
	"github.com/donnyirianto/be-stock-app/modules/auth/domain"
	utils "github.com/donnyirianto/be-stock-app/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *LoginHandler) Login(c *fiber.Ctx) error {
	var requestBody domain.LoginRequest

	// Parse the body
	if err := c.BodyParser(&requestBody); err != nil {
		// Return error response using SendJSONResponseError
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", "Invalid request payload")
	}
	utils.GetLogger().Info("sini")
	// Validate request body
	if err := h.validate.Struct(&requestBody); err != nil {
		// Return error response using SendJSONResponseError
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", err.Error())
	}

	// Call the use case method
	resultLogin, err := h.loginUsecase.Login(&requestBody)
	if err != nil {
		// Return error response using SendJSONResponseError
		return utils.SendJSONResponseError(c, fiber.StatusInternalServerError, "error", err.Error())
	}

	if resultLogin.Status != "success" {
		return utils.SendJSONResponseError(c, resultLogin.Code, resultLogin.Status, resultLogin.Message)
	}

	// Return successful response
	return utils.SendJSONResponse(c, resultLogin.Code, resultLogin.Status, resultLogin.Message, resultLogin.Data)
}
