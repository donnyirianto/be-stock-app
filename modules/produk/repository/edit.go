package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/produk/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *ProdukRepositoryImpl) EditProduk(id string) ([]*domain.ResponseData, error) {

	var respData []*domain.ResponseData

	querySql := `SELECT * from produk WHERE id_produk = ?;`

	err := r.mysqlConn.Raw(querySql, id).Scan(&respData).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return nil, fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return respData, nil

}
