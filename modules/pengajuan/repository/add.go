package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *PengajuanRepositoryImpl) AddPengajuan(req *domain.RequestData, id, username string) error {

	querySql := `INSERT INTO pengajuan (id,subject, keterangan, nama_user, status, created_at, updated_at) VALUES (?, ?, ?, ?, 'PENDING', now(), now())`

	result := r.mysqlConn.Exec(querySql, id, req.Subject, req.Keterangan, username)
	if result.Error != nil {
		utils.GetLogger().Error("Error:", zap.Error(result.Error))
		return fmt.Errorf("Server sedang sibuk, silahkan coba lagi nanti!")
	}

	return nil
}
