package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *UsersRepositoryImpl) DeleteUsers(id string) (string, error) {

	querySql := `DELETE FROM users where username = ?;`

	err := r.mysqlConn.Exec(querySql, id).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return "", fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return "Sukses", nil

}
