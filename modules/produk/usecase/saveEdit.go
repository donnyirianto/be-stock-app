package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/produk/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *produkUsecaseImpl) SaveEditProduk(req *domain.RequestData, id string) (*utils.Response[map[string]any], error) {

	_, err := uc.produkRepository.SaveEditProduk(req, id)
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
