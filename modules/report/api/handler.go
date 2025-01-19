// handler.go

package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/modules/report/usecase"
)

type ReportHandler struct {
	report   usecase.ReportUsecase
	validate *validator.Validate
}

func NewReportHandler(report usecase.ReportUsecase) *ReportHandler {
	return &ReportHandler{
		report:   report,
		validate: validator.New(),
	}
}

func (h *ReportHandler) RegisterRoutes(router fiber.Router) {

	router.Post("/", h.Report)
}
