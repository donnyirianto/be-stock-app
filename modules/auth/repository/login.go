// modules/users/repository/user_repository.go

package repository

import (
	"errors"

	"github.com/donnyirianto/be-stock-app/modules/auth/domain"
	"gorm.io/gorm"
)

// Act Login
func (r *LoginRepositoryImpl) ActLogin(username string) (*domain.HasilLogin, error) {

	var resLogin domain.HasilLogin
	query := `SELECT 
				username,nama,password,id_role
			FROM users 
			WHERE 
				username = ? 
				AND aktif = 'Y';`

	err := r.mysqlConn.Raw(query, username).Scan(&resLogin).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &resLogin, nil
}
