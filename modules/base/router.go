package base

import (
	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/base/api"
	"github.com/donnyirianto/be-stock-app/modules/base/repository"
	"github.com/donnyirianto/be-stock-app/modules/base/usecase"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, mysqlConn *gorm.DB, cfg *config.Config) {

	baseRouter := app.Group("/api/v1/base")
	baseRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Custom response format for JWT errors
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    fiber.StatusUnauthorized,
					"status":  "error",
					"message": "Invalid or expired JWT",
				})
			}
			return nil
		},
	}))
	baseRepository := repository.NewBaseRepository(mysqlConn)
	baseUsecase := usecase.NewBaseUsecase(baseRepository)
	apiHandler := api.NewBaseHandler(baseUsecase)
	apiHandler.RegisterRoutes(baseRouter)
}
