// handler.go

package api

import (
	"github.com/donnyirianto/be-stock-app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *BaseHandler) GetMenu(c *fiber.Ctx) error {

	ctx := c.Context()

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	result, err := h.baseUsecase.GetMenu(ctx, role)
	if err != nil {
		return utils.SendJSONResponseError(c, fiber.StatusInternalServerError, "error", err.Error())
	}

	return utils.SendJSONResponse(c, result.Code, result.Status, result.Message, result.Data)
}
