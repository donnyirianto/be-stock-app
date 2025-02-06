package usecase

import (
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *menuUsecaseImpl) DeleteMenu(id int) (*utils.Response[map[string]interface{}], error) {

	_, err := uc.menuRepository.DeleteMenu(id)
	if err != nil {
		return nil, err
	}

	return &utils.Response[map[string]interface{}]{
		Code:    200,
		Status:  "success",
		Message: "Success",
		Data:    map[string]interface{}{},
	}, nil

}
