package repository

import (
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/produk/domain"
)

type ProdukRepositoryImpl struct {
	cfg       *config.Config
	mysqlConn *gorm.DB
}

func NewProdukRepository(cfg *config.Config, mysqlConn *gorm.DB) domain.ProdukRepository {
	return &ProdukRepositoryImpl{
		cfg:       cfg,
		mysqlConn: mysqlConn,
	}
}
