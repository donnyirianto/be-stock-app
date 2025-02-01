// handler.go

package api

import (
	"github.com/donnyirianto/be-stock-app/modules/base/usecase"

	"github.com/gofiber/fiber/v2"
)

type BaseHandler struct {
	baseUsecase usecase.BaseUsecase
}

func NewBaseHandler(baseUsecase usecase.BaseUsecase) *BaseHandler {
	return &BaseHandler{
		baseUsecase: baseUsecase,
	}
}

func (h *BaseHandler) RegisterRoutes(router fiber.Router) {

	router.Get("/menu", h.GetMenu)
}
