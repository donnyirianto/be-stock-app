package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/produk/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *ProdukRepositoryImpl) SaveEditProduk(req *domain.RequestData, id string) (string, error) {

	querySql := `UPDATE produk SET id_produk = ?, nama= ? ,merk = ?,satuan = ?, harga = ? where id_produk = ?;`

	err := r.mysqlConn.Exec(querySql, req.IdProduk, req.Nama, req.Merk, req.Satuan, req.Harga, id).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return "", fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return "Sukses", nil

}
