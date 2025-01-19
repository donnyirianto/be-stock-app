package repository

import (
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/report/domain"
)

type ReportRepositoryImpl struct {
	cfg       *config.Config
	mysqlConn *gorm.DB
}

func NewReportRepository(cfg *config.Config, mysqlConn *gorm.DB) domain.ReportRepository {
	return &ReportRepositoryImpl{
		cfg:       cfg,
		mysqlConn: mysqlConn,
	}
}
