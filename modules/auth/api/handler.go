package api

import (
	"github.com/donnyirianto/be-stock-app/modules/auth/domain"
	"github.com/donnyirianto/be-stock-app/modules/auth/usecase"
	utils "github.com/donnyirianto/be-stock-app/utils"
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
	router.Post("/login", h.actLogin)
}

func (h *LoginHandler) actLogin(c *fiber.Ctx) error {
	var requestBody domain.LoginRequest

	// Parse the body
	if err := c.BodyParser(&requestBody); err != nil {
		// Return error response using SendJSONResponseError
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", "Invalid request payload")
	}

	// Validate request body
	if err := h.validate.Struct(&requestBody); err != nil {
		// Return error response using SendJSONResponseError
		return utils.SendJSONResponseError(c, fiber.StatusBadRequest, "error", err.Error())
	}

	// Call the use case method
	resultLogin, err := h.loginUsecase.ActLogin(&requestBody)
	if err != nil {
		// Return error response using SendJSONResponseError
		return utils.SendJSONResponseError(c, fiber.StatusInternalServerError, "error", err.Error())
	}

	if resultLogin.Status != "success" {
		return utils.SendJSONResponseError(c, resultLogin.Code, resultLogin.Status, resultLogin.Message)
	}

	// Return successful response
	return utils.SendJSONResponse(c, resultLogin.Code, resultLogin.Status, resultLogin.Message, resultLogin.Result)
}
