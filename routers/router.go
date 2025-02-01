package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	auth "github.com/donnyirianto/be-stock-app/modules/auth"
	base "github.com/donnyirianto/be-stock-app/modules/base"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config, mysqlConn *gorm.DB) {

	auth.RegisterRoutes(app, mysqlConn, cfg) // Auth
	base.RegisterRoutes(app, mysqlConn, cfg) // Base
}
