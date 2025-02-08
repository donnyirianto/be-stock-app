// handler.go

package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/modules/users/usecase"
)

type UsersHandler struct {
	users    usecase.UsersUsecase
	validate *validator.Validate
}

func NewUsersHandler(users usecase.UsersUsecase) *UsersHandler {
	return &UsersHandler{
		users:    users,
		validate: validator.New(),
	}
}

func (h *UsersHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetUsers)
	router.Post("/", h.AddUsers)
	router.Get("/edit/:id", h.EditUsers)
	router.Post("/edit/:id", h.SaveEditUsers)
	router.Delete("/:id", h.DeleteUsers)
}
