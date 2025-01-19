package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config, mysqlConn *gorm.DB) {

	//reportSo.RegisterRoutesV3(app, mysqlConn, cfg) // Report SO
}
