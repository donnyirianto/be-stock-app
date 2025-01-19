// modules/users/usecase/user_usecase.go

package usecase

import (
	"time"

	"github.com/donnyirianto/be-stock-app/modules/refresh/domain"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshUsecase interface {
	ActRefresh(username, nama string) (*domain.HasilRefresh, error)
}

type RefreshUsecaseImpl struct {
	jwtSecret                  string
	jwtSecretRefresh           string
	tokenExpiryDuration        time.Duration
	tokenRefreshExpiryDuration time.Duration
}

func NewRefreshUsecase(jwtSecret string, jwtSecretRefresh string, tokenExpiryDuration time.Duration, tokenRefreshExpiryDuration time.Duration) *RefreshUsecaseImpl {
	return &RefreshUsecaseImpl{
		jwtSecret:                  jwtSecret,
		jwtSecretRefresh:           jwtSecretRefresh,
		tokenExpiryDuration:        tokenExpiryDuration,
		tokenRefreshExpiryDuration: tokenRefreshExpiryDuration,
	}
}

func (uc *RefreshUsecaseImpl) ActRefresh(username, nama string) (*domain.HasilRefresh, error) {

	claims := jwt.MapClaims{
		"username": username,
		"nama":     nama,
		"exp":      time.Now().Add(time.Hour * uc.tokenExpiryDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(uc.jwtSecret))
	if err != nil {
		return nil, domain.ErrGenerateToken
	}

	claimsRefreshToken := jwt.MapClaims{
		"username": username,
		"nama":     nama,
		"exp":      time.Now().Add(time.Hour * uc.tokenRefreshExpiryDuration).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)

	refreshTokenString, err := refreshToken.SignedString([]byte(uc.jwtSecretRefresh))
	if err != nil {
		return nil, domain.ErrGenerateToken
	}

	HasilRefresh := &domain.HasilRefresh{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}

	return HasilRefresh, nil
}
