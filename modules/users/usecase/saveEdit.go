package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/users/domain"
	"github.com/donnyirianto/be-stock-app/utils"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usersUsecaseImpl) SaveEditUsers(req *domain.RequestData, id string) (*utils.Response[map[string]any], error) {

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		req.Password = string(hashedPassword) // Ganti password dengan hash
	}

	_, err := uc.usersRepository.SaveEditUsers(req, id)
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
