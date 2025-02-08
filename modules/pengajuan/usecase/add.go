package usecase

import (
	"fmt"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *pengajuanUsecaseImpl) AddPengajuan(req *domain.RequestData, username string) (*utils.Response[map[string]any], error) {
	id := utils.GenerateID()

	err := uc.pengajuanRepository.AddPengajuan(req, id, username)
	if err != nil {
		return nil, err
	}

	values := []string{}
	for _, p := range req.DetailItem {
		values = append(values, fmt.Sprintf("('%s', '%s','%s',now(),now())", id, p.IdProduk, p.Harga))
	}

	err = uc.pengajuanRepository.AddPengajuanDetail(values)
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
