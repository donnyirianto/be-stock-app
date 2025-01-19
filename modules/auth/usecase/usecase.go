package usecase

import (
	"time"

	"github.com/donnyirianto/be-stock-app/modules/auth/domain"
	"github.com/donnyirianto/be-stock-app/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase interface {
	ActLogin(*domain.LoginRequest) (*utils.Response[map[string]interface{}], error)
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

// ActLogin handles the login logic
func (uc *LoginUsecaseImpl) ActLogin(req *domain.LoginRequest) (*utils.Response[map[string]interface{}], error) {

	repoLogin, err := uc.loginRepository.ActLogin(req.Username)
	if err != nil {
		return nil, err
	}

	if repoLogin == nil {
		return &utils.Response[map[string]interface{}]{
			Code:    404,
			Status:  "error",
			Message: "Username belum terdaftar / terblokir",
			Result:  nil,
		}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(repoLogin.Password), []byte(req.Password)); err != nil {
		return &utils.Response[map[string]interface{}]{
			Code:    401,
			Status:  "error",
			Message: "Invalid username or password",
			Result:  nil,
		}, nil
	}

	claims := jwt.MapClaims{
		"username": req.Username,
		"nama":     repoLogin.Nama,
		"role":     repoLogin.IdRole,
		"exp":      time.Now().Add(time.Hour * uc.tokenExpiryDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(uc.jwtSecret))
	if err != nil {
		return nil, err
	}

	// Generate JWT refresh token
	claimsRefreshToken := jwt.MapClaims{
		"username": req.Username,
		"nama":     repoLogin.Nama,
		"role":     repoLogin.IdRole,
		"exp":      time.Now().Add(time.Hour * uc.tokenRefreshExpiryDuration).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshTokenString, err := refreshToken.SignedString([]byte(uc.jwtSecretRefresh))
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"username":      req.Username,
		"nama":          repoLogin.Nama,
		"role":          repoLogin.IdRole,
		"token":         tokenString,
		"refresh-token": refreshTokenString,
	}

	return &utils.Response[map[string]interface{}]{
		Code:    200,
		Status:  "success",
		Message: "Login successful",
		Result:  result,
	}, nil
}
