package usecase

import (
	"context"

	"github.com/donnyirianto/be-stock-app/modules/base/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

func (uc *BaseUsecaseImpl) GetMenu(ctx context.Context, role string) (*utils.Response[map[string]interface{}], error) {

	var resBase []*domain.BaseItem
	var formatBase []*domain.NewBaseItem

	resBase, err := uc.baseRepository.GetMenu(role)
	if err != nil {
		return nil, err
	}

	formatBase, err = uc.TreeBase(resBase)
	if err != nil {
		return nil, err
	}

	return &utils.Response[map[string]interface{}]{
		Code:    200,
		Status:  "success",
		Message: "Get Data Menu - Success",
		Data:    map[string]interface{}{"menu": formatBase},
	}, nil

}
