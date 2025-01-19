// handler.go

package api

import (
	"errors"

	"github.com/donnyirianto/be-stock-app/modules/refresh/domain"
	"github.com/donnyirianto/be-stock-app/modules/refresh/usecase"
	"github.com/donnyirianto/be-stock-app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshHandler struct {
	refreshUsecase usecase.RefreshUsecase
}
type RefreshRequest struct {
	Token_refresh string `json:"token_refresh"`
}

func NewRefreshHandler(refreshUsecase usecase.RefreshUsecase) *RefreshHandler {
	return &RefreshHandler{refreshUsecase: refreshUsecase}
}

func (h *RefreshHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.actRefresh)
}

func (h *RefreshHandler) actRefresh(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	nama := claims["nama"].(string)

	resultRefresh, err := h.refreshUsecase.ActRefresh(username, nama)
	if err != nil {

		if errors.Is(err, domain.ErrInvalidToken) {

			return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
		}
		return err
	}

	return utils.SendJSONResponse(c, fiber.StatusOK, "success", "Refresh Token Sukses", resultRefresh)
}
