package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *PengajuanRepositoryImpl) SaveEditPengajuan(req *domain.RequestData, id string) error {

	querySql := `UPDATE pengajuan SET
				subject = ?,
				keterangan = ?,
				updated_at = now()
				where id = ?`

	result := r.mysqlConn.Exec(querySql, req.Subject, req.Keterangan, id)
	if result.Error != nil {
		utils.GetLogger().Error("Error:", zap.Error(result.Error))
		return fmt.Errorf("Server sedang sibuk, silahkan coba lagi nanti!")
	}

	return nil
}
