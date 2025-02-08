package master

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/pengajuan/api"
	"github.com/donnyirianto/be-stock-app/modules/pengajuan/repository"
	"github.com/donnyirianto/be-stock-app/modules/pengajuan/usecase"
)

func RegisterRoutes(app *fiber.App, mysqlConn *gorm.DB, cfg *config.Config) {

	pengajuanRouter := app.Group("/api/v1/pengajuan")
	pengajuanRouter.Use(jwtware.New(jwtware.Config{
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
	pengajuanRepository := repository.NewPengajuanRepository(cfg, mysqlConn)
	pengajuanUsecase := usecase.NewPengajuanUsecase(pengajuanRepository)
	apiHandler := api.NewPengajuanHandler(pengajuanUsecase)
	apiHandler.RegisterRoutes(pengajuanRouter)
}
