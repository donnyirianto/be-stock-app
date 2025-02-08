package domain

import (
	"errors"
)

type DetailItem struct {
	IdProduk string `json:"id_produk" validate:"required"`
	Harga    string `json:"harga" validate:"required"`
}
type RequestData struct {
	Subject    string       `json:"subject" validate:"required"`
	Keterangan string       `json:"keterangan" validate:"required"`
	DetailItem []DetailItem `json:"detail_item" validate:"required,dive"`
}

type ResponseData struct {
	Id         string `json:"id"`
	Subject    string `json:"subject"`
	Keterangan string `json:"keterangan"`
	TotalItem  string `json:"total_item"`
	Status     string `json:"status"`
	NamaUser   string `json:"nama_user"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PengajuanRepository interface {
	GetPengajuan() ([]*ResponseData, error)
	AddPengajuan(req *RequestData, id, username string) error
	AddPengajuanDetail(data []string) error
	//EditPengajuan(id string) ([]*ResponseData, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
