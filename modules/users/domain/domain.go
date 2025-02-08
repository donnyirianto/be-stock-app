package domain

import (
	"errors"
)

type RequestData struct {
	Username string `json:"username" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	IdRole   string `json:"id_role" validate:"required"`
	Password string `json:"password"`
	Aktif    string `json:"aktif" validate:"required"`
}

type ResponseData struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	IdRole   string `json:"id_role"`
	NamaRole string `json:"nama_role"`
	Aktif    string `json:"aktif"`
}

type UsersRepository interface {
	GetUsers() ([]*ResponseData, error)
	AddUsers(req *RequestData) (string, error)
	EditUsers(id string) ([]*ResponseData, error)
	SaveEditUsers(req *RequestData, id string) (string, error)
	DeleteUsers(id string) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
