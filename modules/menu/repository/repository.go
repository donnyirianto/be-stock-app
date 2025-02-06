package repository

import (
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/menu/domain"
)

type MenuRepositoryImpl struct {
	cfg       *config.Config
	mysqlConn *gorm.DB
}

func NewMenuRepository(cfg *config.Config, mysqlConn *gorm.DB) domain.MenuRepository {
	return &MenuRepositoryImpl{
		cfg:       cfg,
		mysqlConn: mysqlConn,
	}
}
