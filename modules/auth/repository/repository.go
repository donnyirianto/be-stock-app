// modules/users/repository/user_repository.go

package repository

import (
	"github.com/donnyirianto/be-stock-app/modules/auth/domain"
	"gorm.io/gorm"
)

type LoginRepositoryImpl struct {
	mysqlConn *gorm.DB // GORM database connection
}

func NewLoginRepository(mysqlConn *gorm.DB) domain.LoginRepository {
	return &LoginRepositoryImpl{
		mysqlConn: mysqlConn,
	}
}
