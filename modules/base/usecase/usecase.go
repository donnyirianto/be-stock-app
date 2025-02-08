package usecase

import (
	"context"

	"github.com/donnyirianto/be-stock-app/modules/base/domain"
	"github.com/donnyirianto/be-stock-app/utils"
)

type BaseUsecase interface {
	GetMenu(ctx context.Context, role string) (*utils.Response[map[string]any], error)
	TreeBase(baseItems []*domain.BaseItem) ([]*domain.NewBaseItem, error)
}

type BaseUsecaseImpl struct {
	baseRepository domain.BaseRepository
}

func NewBaseUsecase(baseRepository domain.BaseRepository) *BaseUsecaseImpl {
	return &BaseUsecaseImpl{
		baseRepository: baseRepository,
	}
}
