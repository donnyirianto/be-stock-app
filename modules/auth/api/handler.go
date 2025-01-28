package api

import (
	"github.com/donnyirianto/be-stock-app/modules/auth/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
	loginUsecase usecase.LoginUsecase
	validate     *validator.Validate
}

func NewLoginHandler(loginUsecase usecase.LoginUsecase) *LoginHandler {
	return &LoginHandler{
		loginUsecase: loginUsecase,
		validate:     validator.New(),
	}
}

// RegisterRoutes registers user-related routes
func (h *LoginHandler) RegisterRoutes(router fiber.Router) {
	// Define user-related routes here
	router.Post("/login", h.Login)
	router.Get("/", h.RefreshToken)
}
