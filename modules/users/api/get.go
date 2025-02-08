package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/utils"
)

func (h *UsersHandler) GetUsers(c *fiber.Ctx) error {

	resp, err := h.users.GetUsers()
	if err != nil {
		return utils.SendJSONResponseError(c, resp.Code, resp.Status, resp.Message)
	}

	return utils.SendJSONResponse(c, resp.Code, resp.Status, resp.Message, resp.Data)
}
