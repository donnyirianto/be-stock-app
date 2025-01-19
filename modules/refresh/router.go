package refresh

import (
	"time"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/refresh/api"
	"github.com/donnyirianto/be-stock-app/modules/refresh/usecase"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutesV1(app *fiber.App, db *gorm.DB, cfg *config.Config) {

	refreshRouter := app.Group("/api/v1/refresh-token")
	refreshRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWTSecretRefresh)},
	}))
	refreshUsecase := usecase.NewRefreshUsecase(cfg.JWTSecret, cfg.JWTSecretRefresh, time.Duration(cfg.TokenExpiryDuration), time.Duration(cfg.TokenRefreshExpiryDuration))
	apiHandler := api.NewRefreshHandler(refreshUsecase)
	apiHandler.RegisterRoutes(refreshRouter)
}
