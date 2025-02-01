package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"

	"github.com/donnyirianto/be-stock-app/repository"
	"github.com/donnyirianto/be-stock-app/routers"
	"github.com/donnyirianto/be-stock-app/utils"
)

func main() {

	pathConfig := flag.String("config", "./", "Path Config")
	flag.Parse()

	// Initialize logger
	utils.InitLogger()
	defer utils.GetLogger().Sync()

	// Load configuration
	cfg, err := config.LoadConfig(*pathConfig)
	if err != nil {
		utils.GetLogger().Fatal("Failed to load config", zap.Error(err))
	}

	// Connect to Mysql
	mysqlConn, err := repository.ConnectMySQL(cfg)
	if err != nil {
		utils.GetLogger().Fatal("Failed to connect to Database", zap.Error(err))
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ServerHeader:            "Stock APP",
		AppName:                 "stock-app",
		JSONEncoder:             sonic.Marshal,
		JSONDecoder:             sonic.Unmarshal,
		Prefork:                 cfg.App.Prefork,
		EnableTrustedProxyCheck: true,
	})

	// Register middleware
	registerMiddleware(app, cfg)

	// Setup routes
	setupRoutes(app, cfg, mysqlConn)

	// Start server
	port := cfg.App.Port
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

}

func registerMiddleware(app *fiber.App, cfg *config.Config) {
	// Generate trace_id

	// Recover middleware
	app.Use(recover.New())

	// Health check
	app.Use(healthcheck.New())

	// Custom error handler middleware
	app.Use(utils.ErrorHandlerMiddleware(utils.GetLogger()))

	// Custom logger middleware using Zap logger
	app.Use(utils.LoggerMiddleware(utils.GetLogger(), cfg))

	// Compression middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// CORS middleware
	app.Use(cors.New())

	//Rate Limit
	app.Use(limiter.New(limiter.Config{
		Max:        2000,
		Expiration: 5 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"code":    429,
				"status":  "error",
				"message": "Server sedang sibuk, silahkan dicoba beberapa saat lagi!",
			})
		},
	}))
}

func setupRoutes(app *fiber.App, cfg *config.Config, mysqlConn *gorm.DB) {

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"code":    200,
			"status":  "success",
			"message": "service running",
		})
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	api.Get("/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"code":    200,
			"status":  "success",
			"message": "service running",
		})
	})

	routers.RegisterRoutes(app, cfg, mysqlConn)

}
