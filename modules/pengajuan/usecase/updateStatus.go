package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *pengajuanUsecaseImpl) UpdateStatus(req *domain.RequestDataStatus, id string) (*utils.Response[map[string]any], error) {

	err := uc.pengajuanRepository.UpdateStatus(req, id)
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
