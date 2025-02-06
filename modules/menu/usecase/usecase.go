package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/menu/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

type MenuUsecase interface {
	GetMenu() (*utils.Response[map[string]interface{}], error)
	AddMenu(*domain.RequestData) (*utils.Response[map[string]interface{}], error)
	EditMenu(id int) (*utils.Response[map[string]interface{}], error)
	DeleteMenu(id int) (*utils.Response[map[string]interface{}], error)
}

type menuUsecaseImpl struct {
	menuRepository domain.MenuRepository
}

func NewMenuUsecase(menuRepository domain.MenuRepository) *menuUsecaseImpl {
	return &menuUsecaseImpl{
		menuRepository: menuRepository,
	}
}
