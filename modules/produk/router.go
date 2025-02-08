package master

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/produk/api"
	"github.com/donnyirianto/be-stock-app/modules/produk/repository"
	"github.com/donnyirianto/be-stock-app/modules/produk/usecase"
)

func RegisterRoutes(app *fiber.App, mysqlConn *gorm.DB, cfg *config.Config) {

	produkRouter := app.Group("/api/v1/produk")
	produkRouter.Use(jwtware.New(jwtware.Config{
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
	produkRepository := repository.NewProdukRepository(cfg, mysqlConn)
	produkUsecase := usecase.NewProdukUsecase(produkRepository)
	apiHandler := api.NewProdukHandler(produkUsecase)
	apiHandler.RegisterRoutes(produkRouter)
}
