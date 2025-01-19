package master

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/report/api"
	"github.com/donnyirianto/be-stock-app/modules/report/repository"
	"github.com/donnyirianto/be-stock-app/modules/report/usecase"
)

func RegisterRoutesV3(app *fiber.App, mysqlConn *gorm.DB, cfg *config.Config) {

	reportRouter := app.Group("/api/v1/sokas/report")
	reportRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Custom response format for JWT errors
			if err != nil {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"code":    fiber.StatusForbidden,
					"status":  "error",
					"message": "Invalid or expired JWT",
				})
			}
			return nil
		},
	}))
	reportRepository := repository.NewReportRepository(cfg, mysqlConn)
	reportUsecase := usecase.NewReportUsecase(reportRepository)
	apiHandler := api.NewReportHandler(reportUsecase)
	apiHandler.RegisterRoutes(reportRouter)
}
