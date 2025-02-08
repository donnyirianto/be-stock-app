package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/users/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *usersUsecaseImpl) AddUsers(req *domain.RequestData) (*utils.Response[map[string]any], error) {

	_, err := uc.usersRepository.AddUsers(req)
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
