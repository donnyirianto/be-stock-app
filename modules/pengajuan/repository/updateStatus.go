package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *PengajuanRepositoryImpl) UpdateStatus(req *domain.RequestDataStatus, id string) error {

	querySql := `UPDATE pengajuan SET
				status = UPPER(?),
				updated_at = now()
				where id = ?`

	result := r.mysqlConn.Exec(querySql, req.Status, id)
	if result.Error != nil {
		utils.GetLogger().Error("Error:", zap.Error(result.Error))
		return fmt.Errorf("Server sedang sibuk, silahkan coba lagi nanti!")
	}

	return nil
}
