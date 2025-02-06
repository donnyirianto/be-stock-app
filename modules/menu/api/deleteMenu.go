package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/utils"
)

func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", "Payload tidak sesuai!")
	}

	resp, err := h.menu.DeleteMenu(id)
	if err != nil {
		return utils.SendJSONResponseError(c, resp.Code, resp.Status, resp.Message)
	}

	return utils.SendJSONResponse(c, resp.Code, resp.Status, resp.Message, resp.Data)
}
