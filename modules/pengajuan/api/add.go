package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (h *PengajuanHandler) AddPengajuan(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	var requestBody domain.RequestData

	if err := c.BodyParser(&requestBody); err != nil {
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", "Payload tidak sesuai!")
	}

	if err := h.validate.Struct(&requestBody); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("failed on the '%s' tag", err.Tag())
		}
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", validationErrors)
	}

	resp, err := h.pengajuan.AddPengajuan(&requestBody, username)
	if err != nil {
		return utils.SendJSONResponseError(c, resp.Code, resp.Status, resp.Message)
	}

	return utils.SendJSONResponse(c, resp.Code, resp.Status, resp.Message, resp.Data)
}
