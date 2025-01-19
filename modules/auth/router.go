package login

import (
	"time"

	config "github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/auth/api"
	"github.com/donnyirianto/be-stock-app/modules/auth/repository"
	"github.com/donnyirianto/be-stock-app/modules/auth/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutesV1(app *fiber.App, chConn *gorm.DB, cfg *config.Config) {

	loginRouter := app.Group("/api/v1/auth")
	loginRepository := repository.NewLoginRepository(chConn)
	loginUsecase := usecase.NewLoginUsecase(loginRepository, cfg.JWTSecret, cfg.JWTSecretRefresh, time.Duration(cfg.TokenExpiryDuration), time.Duration(cfg.TokenRefreshExpiryDuration))
	apiHandler := api.NewLoginHandler(loginUsecase)
	apiHandler.RegisterRoutes(loginRouter)
}
