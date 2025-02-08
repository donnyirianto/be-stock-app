package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/users/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

type UsersUsecase interface {
	GetUsers() (*utils.Response[map[string]any], error)
	AddUsers(*domain.RequestData) (*utils.Response[map[string]any], error)
	EditUsers(id string) (*utils.Response[map[string]any], error)
	SaveEditUsers(req *domain.RequestData, id string) (*utils.Response[map[string]any], error)
	DeleteUsers(id string) (*utils.Response[map[string]any], error)
}

type usersUsecaseImpl struct {
	usersRepository domain.UsersRepository
}

func NewUsersUsecase(usersRepository domain.UsersRepository) *usersUsecaseImpl {
	return &usersUsecaseImpl{
		usersRepository: usersRepository,
	}
}
