package usecase

import "github.com/donnyirianto/be-stock-app/utils"

func (uc *produkUsecaseImpl) GetProduk() (*utils.Response[map[string]any], error) {

	resp, err := uc.produkRepository.GetProduk()
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return &utils.Response[map[string]any]{
			Code:    200,
			Status:  "not found",
			Message: "Data Not Found",
			Data:    map[string]any{"produk": []map[string]any{}},
		}, nil
	}
	return &utils.Response[map[string]any]{
		Code:    200,
		Status:  "success",
		Message: "Get Data Produk - Success",
		Data:    map[string]any{"produk": resp},
	}, nil

}
