package usecase

import (
	"context"

	"github.com/donnyirianto/be-stock-app/modules/report/domain"
)

func (uc *reportUsecaseImpl) Report(ctx context.Context, username string, req *domain.RequestData) (any, error) {

	resp, err := uc.reportRepository.Report(ctx, req, []string{})
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		respData := []map[string]interface{}{}
		return respData, nil
	}

	return resp, nil
}
