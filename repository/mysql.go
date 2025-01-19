package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/utils"
)

func ConnectMySQL(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=False&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1000 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		if err := sqlDB.Close(); err != nil {
			utils.GetLogger().Fatal("Failed to close database connection", zap.Error(err))
		}
		return nil, fmt.Errorf("failed to get SQL DB from GORM: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.Database.DbMaxIdle)
	sqlDB.SetMaxOpenConns(cfg.Database.DbMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.DbMaxIdle) * time.Minute)

	return db, nil
}
