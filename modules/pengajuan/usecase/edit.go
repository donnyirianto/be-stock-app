package usecase

import "github.com/donnyirianto/be-stock-app/utils"

func (uc *pengajuanUsecaseImpl) EditPengajuan(id string) (*utils.Response[map[string]any], error) {

	resp, err := uc.pengajuanRepository.EditPengajuan(id)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return &utils.Response[map[string]any]{
			Code:    200,
			Status:  "not found",
			Message: "Data Not Found",
			Data:    map[string]any{"pengajuan": map[string]any{}},
		}, nil
	}

	respDetail, err := uc.pengajuanRepository.EditPengajuanDetail(id)
	if err != nil {
		return nil, err
	}

	return &utils.Response[map[string]any]{
		Code:    200,
		Status:  "success",
		Message: "Get Data Pengajuan - Success",
		Data:    map[string]any{"pengajuan": resp[0], "detail_item": respDetail},
	}, nil

}
