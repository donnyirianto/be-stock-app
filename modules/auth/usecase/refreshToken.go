package usecase

import (
	"time"

	"github.com/donnyirianto/be-stock-app/utils"

	"github.com/golang-jwt/jwt/v5"
)

// ActLogin handles the login logic
func (uc *LoginUsecaseImpl) RefreshToken(username, nama, role string) (*utils.Response[map[string]interface{}], error) {

	claims := jwt.MapClaims{
		"username": username,
		"nama":     nama,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * uc.tokenExpiryDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(uc.jwtSecret))
	if err != nil {
		return nil, err
	}

	claimsRefreshToken := jwt.MapClaims{
		"username": username,
		"nama":     nama,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * uc.tokenRefreshExpiryDuration).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)

	refreshTokenString, err := refreshToken.SignedString([]byte(uc.jwtSecretRefresh))
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"accessToken":  tokenString,
		"refreshToken": refreshTokenString,
	}

	return &utils.Response[map[string]interface{}]{
		Code:    200,
		Status:  "success",
		Message: "Generate Refresh Token successful",
		Data:    result,
	}, nil
}
