package usecase

import "github.com/donnyirianto/be-stock-app/utils"

func (uc *pengajuanUsecaseImpl) GetPengajuan() (*utils.Response[map[string]any], error) {

	resp, err := uc.pengajuanRepository.GetPengajuan()
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return &utils.Response[map[string]any]{
			Code:    200,
			Status:  "not found",
			Message: "Data Not Found",
			Data:    map[string]any{"pengajuan": []map[string]any{}},
		}, nil
	}
	return &utils.Response[map[string]any]{
		Code:    200,
		Status:  "success",
		Message: "Get Data Pengajuan - Success",
		Data:    map[string]any{"pengajuan": resp},
	}, nil

}
