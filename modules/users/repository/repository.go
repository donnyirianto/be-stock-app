package repository

import (
	"gorm.io/gorm"

	"github.com/donnyirianto/be-stock-app/config"
	"github.com/donnyirianto/be-stock-app/modules/users/domain"
)

type UsersRepositoryImpl struct {
	cfg       *config.Config
	mysqlConn *gorm.DB
}

func NewUsersRepository(cfg *config.Config, mysqlConn *gorm.DB) domain.UsersRepository {
	return &UsersRepositoryImpl{
		cfg:       cfg,
		mysqlConn: mysqlConn,
	}
}
