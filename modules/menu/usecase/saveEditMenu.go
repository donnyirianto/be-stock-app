package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/menu/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *menuUsecaseImpl) SaveEditMenu(req *domain.RequestData, id int) (*utils.Response[map[string]any], error) {

	_, err := uc.menuRepository.SaveEditMenu(req, id)
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
