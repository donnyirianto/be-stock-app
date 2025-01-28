package api

import (
	utils "github.com/donnyirianto/be-stock-app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *LoginHandler) RefreshToken(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	nama := claims["nama"].(string)

	resultRefresh, err := h.loginUsecase.RefreshToken(username, nama)
	if err != nil {

		return utils.SendJSONResponseError(c, fiber.StatusUnauthorized, "error", err.Error())

	}

	if resultRefresh.Status != "success" {
		return utils.SendJSONResponseError(c, resultRefresh.Code, resultRefresh.Status, resultRefresh.Message)
	}

	// Return successful response
	return utils.SendJSONResponse(c, resultRefresh.Code, resultRefresh.Status, resultRefresh.Message, resultRefresh.Data)
}
