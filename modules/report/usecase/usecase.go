package usecase

import (
	"context"

	"github.com/donnyirianto/be-stock-app/modules/report/domain"
)

type ReportUsecase interface {
	Report(ctx context.Context, username string, dataRequest *domain.RequestData) (any, error)
}

type reportUsecaseImpl struct {
	reportRepository domain.ReportRepository
}

func NewReportUsecase(reportRepository domain.ReportRepository) *reportUsecaseImpl {
	return &reportUsecaseImpl{
		reportRepository: reportRepository,
	}
}
