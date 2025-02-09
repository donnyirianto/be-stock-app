package usecase

import (
	"github.com/donnyirianto/be-stock-app/modules/pengajuan/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

type PengajuanUsecase interface {
	GetPengajuan() (*utils.Response[map[string]any], error)
	AddPengajuan(req *domain.RequestData, usernames string) (*utils.Response[map[string]any], error)
	EditPengajuan(id string) (*utils.Response[map[string]any], error)
	SaveEditPengajuan(req *domain.RequestData, id string) (*utils.Response[map[string]any], error)
	UpdateStatus(req *domain.RequestDataStatus, id string) (*utils.Response[map[string]any], error)
}

type pengajuanUsecaseImpl struct {
	pengajuanRepository domain.PengajuanRepository
}

func NewPengajuanUsecase(pengajuanRepository domain.PengajuanRepository) *pengajuanUsecaseImpl {
	return &pengajuanUsecaseImpl{
		pengajuanRepository: pengajuanRepository,
	}
}
