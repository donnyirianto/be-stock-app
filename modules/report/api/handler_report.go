package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/donnyirianto/be-stock-app/modules/report/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (h *ReportHandler) Report(c *fiber.Ctx) error {
	ctx := c.Context()

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userDetails := claims["user"].(map[string]any)
	username := userDetails["username"].(string)

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

	dataSales, err := h.report.Report(ctx, username, &requestBody)
	if err != nil {
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", err.Error())
	}

	return utils.SendJSONResponse(c, fiber.StatusOK, "success", "Read Data Report - Sukses", dataSales)
}
