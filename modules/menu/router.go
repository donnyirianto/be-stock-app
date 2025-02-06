package master

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/menu/api"
	"github.com/donnyirianto/be-stock-app/modules/menu/repository"
	"github.com/donnyirianto/be-stock-app/modules/menu/usecase"
)

func RegisterRoutes(app *fiber.App, mysqlConn *gorm.DB, cfg *config.Config) {

	menuRouter := app.Group("/api/v1/settings/menu")
	menuRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Custom response format for JWT errors
			fmt.Println(err)
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
	menuRepository := repository.NewMenuRepository(cfg, mysqlConn)
	menuUsecase := usecase.NewMenuUsecase(menuRepository)
	apiHandler := api.NewMenuHandler(menuUsecase)
	apiHandler.RegisterRoutes(menuRouter)
}
