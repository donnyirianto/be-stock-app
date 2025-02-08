package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *PengajuanRepositoryImpl) GetPengajuan() ([]*domain.ResponseData, error) {

	var respData []*domain.ResponseData

	querySql := `SELECT a.id, c.nama as nama_user,a.subject,a.keterangan,COUNT(*) AS total_item,a.status,a.created_at,a.updated_at FROM pengajuan a 
				LEFT JOIN pengajuan_detail b ON a.id = b.id_pengajuan
				LEFT JOIN users c on a.nama_user = c.username
				GROUP BY a.id
				ORDER BY updated_at desc,created_at desc`

	err := r.mysqlConn.Raw(querySql).Scan(&respData).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return nil, fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return respData, nil

}
