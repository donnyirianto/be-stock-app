package usecase

import (
	"time"

	"github.com/donnyirianto/be-stock-app/modules/auth/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

type LoginUsecase interface {
	Login(*domain.LoginRequest) (*utils.Response[map[string]any], error)
	RefreshToken(username, nama, role string) (*utils.Response[map[string]any], error)
}

type LoginUsecaseImpl struct {
	loginRepository            domain.LoginRepository
	jwtSecret                  string
	jwtSecretRefresh           string
	tokenExpiryDuration        time.Duration
	tokenRefreshExpiryDuration time.Duration
}

func NewLoginUsecase(loginRepository domain.LoginRepository, jwtSecret string, jwtSecretRefresh string, tokenExpiryDuration time.Duration, tokenRefreshExpiryDuration time.Duration) *LoginUsecaseImpl {
	return &LoginUsecaseImpl{
		loginRepository:            loginRepository,
		jwtSecret:                  jwtSecret,
		jwtSecretRefresh:           jwtSecretRefresh,
		tokenExpiryDuration:        tokenExpiryDuration,
		tokenRefreshExpiryDuration: tokenRefreshExpiryDuration,
	}
}
