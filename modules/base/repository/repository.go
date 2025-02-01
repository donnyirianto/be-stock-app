package repository

import (
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/modules/base/domain"
)

type BaseRepositoryImpl struct {
	mysqlConn *gorm.DB
}

func NewBaseRepository(mysqlConn *gorm.DB) domain.BaseRepository {
	return &BaseRepositoryImpl{
		mysqlConn: mysqlConn,
	}
}
