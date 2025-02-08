package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	auth "github.com/donnyirianto/be-stock-app/modules/auth"
	base "github.com/donnyirianto/be-stock-app/modules/base"
	menu "github.com/donnyirianto/be-stock-app/modules/menu"
	pengajuan "github.com/donnyirianto/be-stock-app/modules/pengajuan"
	produk "github.com/donnyirianto/be-stock-app/modules/produk"
	users "github.com/donnyirianto/be-stock-app/modules/users"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config, mysqlConn *gorm.DB) {

	auth.RegisterRoutes(app, mysqlConn, cfg) // Auth
	base.RegisterRoutes(app, mysqlConn, cfg) // Base

	//Settings
	menu.RegisterRoutes(app, mysqlConn, cfg)  // Menu
	users.RegisterRoutes(app, mysqlConn, cfg) // Users

	//Produk

	produk.RegisterRoutes(app, mysqlConn, cfg)    // Produk
	pengajuan.RegisterRoutes(app, mysqlConn, cfg) // Pengajuan
}
