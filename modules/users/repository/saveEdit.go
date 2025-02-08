package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/users/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *UsersRepositoryImpl) SaveEditUsers(req *domain.RequestData, id int) (string, error) {

	querySql := `UPDATE users set username = ?, password= ? ,nama = ?,aktif = ? WhERE username = ?;`

	err := r.mysqlConn.Exec(querySql, req.Username, req.Password, req.Nama, req.Aktif, id).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return "", fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return "Sukses", nil

}
