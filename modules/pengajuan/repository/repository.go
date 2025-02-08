package repository

import (
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
)

type PengajuanRepositoryImpl struct {
	cfg       *config.Config
	mysqlConn *gorm.DB
}

func NewPengajuanRepository(cfg *config.Config, mysqlConn *gorm.DB) domain.PengajuanRepository {
	return &PengajuanRepositoryImpl{
		cfg:       cfg,
		mysqlConn: mysqlConn,
	}
}
