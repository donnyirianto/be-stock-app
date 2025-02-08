package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/produk/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

type ProdukUsecase interface {
	GetProduk() (*utils.Response[map[string]any], error)
	AddProduk(*domain.RequestData) (*utils.Response[map[string]any], error)
	EditProduk(id string) (*utils.Response[map[string]any], error)
	SaveEditProduk(req *domain.RequestData, id string) (*utils.Response[map[string]any], error)
	DeleteProduk(id string) (*utils.Response[map[string]any], error)
}

type produkUsecaseImpl struct {
	produkRepository domain.ProdukRepository
}

func NewProdukUsecase(produkRepository domain.ProdukRepository) *produkUsecaseImpl {
	return &produkUsecaseImpl{
		produkRepository: produkRepository,
	}
}
