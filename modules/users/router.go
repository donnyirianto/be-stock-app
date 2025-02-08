package master

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/users/api"
	"github.com/donnyirianto/be-stock-app/modules/users/repository"
	"github.com/donnyirianto/be-stock-app/modules/users/usecase"
)

func RegisterRoutes(app *fiber.App, mysqlConn *gorm.DB, cfg *config.Config) {

	usersRouter := app.Group("/api/v1/settings/users")
	usersRouter.Use(jwtware.New(jwtware.Config{
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
	usersRepository := repository.NewUsersRepository(cfg, mysqlConn)
	usersUsecase := usecase.NewUsersUsecase(usersRepository)
	apiHandler := api.NewUsersHandler(usersUsecase)
	apiHandler.RegisterRoutes(usersRouter)
}
