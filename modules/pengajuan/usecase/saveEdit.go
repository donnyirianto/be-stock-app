package usecase

import (
	"fmt"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *pengajuanUsecaseImpl) SaveEditPengajuan(req *domain.RequestData, id string) (*utils.Response[map[string]any], error) {

	err := uc.pengajuanRepository.SaveEditPengajuan(req, id)
	if err != nil {
		return nil, err
	}

	values := []string{}
	for _, p := range req.DetailItem {
		values = append(values, fmt.Sprintf("('%s', '%s','%s',now(),now())", id, p.IdProduk, p.Harga))
	}

	err = uc.pengajuanRepository.SaveEditPengajuanDetail(values, id)
	if err != nil {
		return nil, err
	}

	return &utils.Response[map[string]any]{
		Code:    200,
		Status:  "success",
		Message: "Success",
		Data:    map[string]any{},
	}, nil

}
