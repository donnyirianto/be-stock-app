package domain

import (
	"errors"
)

type RequestData struct {
	Nama   string `json:"nama" validate:"required"`
	Link   string `json:"link" validate:"required"`
	IdMain string `json:"id_main"`
	Urut   string `json:"urut" validate:"required"`
	Active string `json:"active" validate:"required"`
}

type ResponseData struct {
	ID     string `json:"id"`
	Nama   string `json:"nama"`
	Link   string `json:"link"`
	Urut   string `json:"urut"`
	IdMain string `json:"id_main"`
	Active string `json:"active"`
}

type MenuRepository interface {
	GetMenu() ([]*ResponseData, error)
	AddMenu(req *RequestData) (string, error)
	EditMenu(id int) ([]*ResponseData, error)
	SaveEditMenu(req *RequestData, id int) (string, error)
	DeleteMenu(id int) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
