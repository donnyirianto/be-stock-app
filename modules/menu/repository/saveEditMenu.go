package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/menu/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *MenuRepositoryImpl) SaveEditMenu(req *domain.RequestData, id int) (string, error) {

	querySql := `UPDATE m_menu set nama = ?, link= ? ,id_main = ?,urut = ?,active = ? WhERE id = ?;`

	err := r.mysqlConn.Exec(querySql, req.Nama, req.Link, req.IdMain, req.Urut, req.Active, id).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return "", fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return "Sukses", nil

}
