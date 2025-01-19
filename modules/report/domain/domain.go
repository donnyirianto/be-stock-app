package domain

import (
	"context"
	"errors"
)

type RequestData struct {
	TanggalAwal  string `json:"tanggal_awal" validate:"required"`
	TanggalAkhir string `json:"tanggal_akhir" validate:"required"`
}

type ResponseData struct {
	Kdcab        string `json:"kdcab"`
	Toko         string `json:"toko"`
	NamaToko     string `json:"nama_toko"`
	Tanggal      string `json:"tanggal"`
	DataKomputer int    `json:"data_komputer"`
	DataAktual   int    `json:"data_aktual"`
	TotalSelisih int    `json:"total_selisih"`
	FileBa       string `json:"fileba"`
}
type Toko struct {
	Kodetoko string `json:"kodetoko"`
	Namatoko string `json:"namatoko"`
}

type PayloadCache struct {
	Toko []Toko `json:"toko"`
}
type ReportRepository interface {
	Report(ctx context.Context, reqData *RequestData, listToko []string) ([]*ResponseData, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
