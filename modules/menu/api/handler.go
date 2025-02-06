// handler.go

package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/modules/menu/usecase"
)

type MenuHandler struct {
	menu     usecase.MenuUsecase
	validate *validator.Validate
}

func NewMenuHandler(menu usecase.MenuUsecase) *MenuHandler {
	return &MenuHandler{
		menu:     menu,
		validate: validator.New(),
	}
}

func (h *MenuHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetMenu)
	router.Post("/", h.AddMenu)
	router.Get("/edit/:id", h.EditMenu)
	router.Delete("/:id", h.DeleteMenu)
}
