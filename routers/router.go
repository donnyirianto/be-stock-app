package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	auth "github.com/donnyirianto/be-stock-app/modules/auth"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config, mysqlConn *gorm.DB) {

	auth.RegisterRoutes(app, mysqlConn, cfg) // Auth
}
