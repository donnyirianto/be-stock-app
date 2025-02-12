package domain

import (
	"errors"
)

type RequestData struct {
	IdProduk string `json:"id_produk" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Merk     string `json:"merk" validate:"required"`
	Tipe     string `json:"tipe" validate:"required"`
	Satuan   string `json:"satuan" validate:"required"`
	Harga    string `json:"harga" validate:"required"`
}

type ResponseData struct {
	IdProduk string `json:"id_produk"`
	Nama     string `json:"nama"`
	Merk     string `json:"merk"`
	Tipe     string `json:"tipe"`
	Satuan   string `json:"satuan"`
	Harga    string `json:"harga"`
}

type ProdukRepository interface {
	GetProduk() ([]*ResponseData, error)
	AddProduk(req *RequestData) (string, error)
	EditProduk(id string) ([]*ResponseData, error)
	SaveEditProduk(req *RequestData, id string) (string, error)
	DeleteProduk(id string) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
