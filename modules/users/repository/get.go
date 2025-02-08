package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/users/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *UsersRepositoryImpl) GetUsers() ([]*domain.ResponseData, error) {

	var respData []*domain.ResponseData

	querySql := `SELECT * from m_users`

	err := r.mysqlConn.Raw(querySql).Scan(&respData).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return nil, fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return respData, nil

}
