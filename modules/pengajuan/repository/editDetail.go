package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *PengajuanRepositoryImpl) EditPengajuanDetail(id string) ([]*domain.DetailItemPengajuan, error) {

	var respData []*domain.DetailItemPengajuan

	querySql := `SELECT a.id_produk,b.nama,b.merk,b.satuan,cast(a.harga as unsigned) as harga FROM pengajuan_detail a 
				LEFT JOIN produk b ON a.id_produk = b.id_produk
				WHERE a.id_pengajuan = ?;`

	err := r.mysqlConn.Raw(querySql, id).Scan(&respData).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return nil, fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return respData, nil

}
