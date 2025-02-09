package repository

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *PengajuanRepositoryImpl) SaveEditPengajuanDetail(data []string, id string) error {

	err := r.mysqlConn.Exec("DELETE FROM pengajuan_detail where id_pengajuan = ?", id).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	querySql := fmt.Sprintf(`INSERT INTO pengajuan_detail(id_pengajuan,id_produk,harga,created_at,updated_at) values %s;`, strings.Join(data, ", "))

	err = r.mysqlConn.Exec(querySql).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return nil
}
