package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/utils"
)

func (h *ProdukHandler) EditProduk(c *fiber.Ctx) error {

	id := c.Params("id")

	resp, err := h.produk.EditProduk(id)
	if err != nil {
		return utils.SendJSONResponseError(c, resp.Code, resp.Status, resp.Message)
	}

	return utils.SendJSONResponse(c, resp.Code, resp.Status, resp.Message, resp.Data)
}
