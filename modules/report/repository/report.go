package repository

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/donnyirianto/be-stock-app/modules/report/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (r *ReportRepositoryImpl) Report(ctx context.Context, requestData *domain.RequestData, listToko []string) ([]*domain.ResponseData, error) {

	var respData []*domain.ResponseData

	querySql := `SELECT a.kdcab,a.toko, b.nama_toko, a.tanggal, 
					ifnull((sales_kasir_data_komputer + sales_serah_data_komputer + sales_reguler_data_komputer + sales_before_data_komputer + kas_finance + rrak_rencana),0) as data_komputer,
					ifnull(sales_kasir_uang_fisik +sales_serah_uang_fisik +sales_reguler_uang_fisik +sales_before_uang_fisik + kas_uang_modal_kasir + kas_uang_modal_brankas  + kas_tukar_uang + rrak_fisik_dana + rrak_nota,0) as data_aktual,
					total_selisih,ifnull(fileba,'') as file_ba
				FROM so_dana_kas.sdk_so a
				LEFT JOIN db_ori_v3.m_toko b on a.toko = b.kdtk
				WHERE 
				a.tanggal between ? and ?
				AND a.toko in (?)
				AND status_so = 'VALIDASI';`

	err := r.mysqlConn.Raw(querySql, requestData.TanggalAwal, requestData.TanggalAkhir, listToko).Scan(&respData).Error
	if err != nil {
		utils.GetLogger().Error("Error:", zap.Error(err))
		return nil, fmt.Errorf("Server sedang sibuk, silahkan dapat dicoba beberapa saat lagi!")
	}

	return respData, nil

}
