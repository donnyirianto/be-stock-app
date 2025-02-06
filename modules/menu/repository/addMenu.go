package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/menu/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *MenuRepositoryImpl) AddMenu(req *domain.RequestData) (string, error) {

	querySql := `INSERT INTO m_menu set nama = ?, link= ? ,id_main = ?,urut = ?,active = ?;`

	err := r.mysqlConn.Exec(querySql, req.Nama, req.Link, req.IdMain, req.Urut, req.Active).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return "", fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return "Sukses", nil

}
