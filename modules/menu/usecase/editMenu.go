package usecase

import "github.com/donnyirianto/be-stock-app/utils"

func (uc *menuUsecaseImpl) EditMenu(id int) (*utils.Response[map[string]interface{}], error) {

	resp, err := uc.menuRepository.EditMenu(id)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return &utils.Response[map[string]interface{}]{
			Code:    404,
			Status:  "not found",
			Message: "Data Not Found",
			Data:    map[string]interface{}{},
		}, nil
	}

	return &utils.Response[map[string]interface{}]{
		Code:    200,
		Status:  "success",
		Message: "Get Data Menu - Success",
		Data:    map[string]interface{}{"menu": resp[0]},
	}, nil

}
