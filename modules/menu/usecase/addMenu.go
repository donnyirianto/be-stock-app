package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/menu/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *menuUsecaseImpl) AddMenu(req *domain.RequestData) (*utils.Response[map[string]any], error) {

	_, err := uc.menuRepository.AddMenu(req)
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
