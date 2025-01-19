package domain

import "errors"

type Refresh struct {
	TokenRefresh string `json:"token_refresh"`
}

type HasilRefresh struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRepository interface {
	ActRefresh(paramsRefresh *Refresh) (string, error)
}

var ErrInvalidToken = errors.New("invalid Token")
var ErrGenerateToken = errors.New("Internal Server Error")
