// handler.go

package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/modules/pengajuan/usecase"
)

type PengajuanHandler struct {
	pengajuan usecase.PengajuanUsecase
	validate  *validator.Validate
}

func NewPengajuanHandler(pengajuan usecase.PengajuanUsecase) *PengajuanHandler {
	return &PengajuanHandler{
		pengajuan: pengajuan,
		validate:  validator.New(),
	}
}

func (h *PengajuanHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetPengajuan)
	router.Post("/", h.AddPengajuan)
	router.Get("/edit/:id", h.EditPengajuan)
	router.Post("/edit/:id", h.SaveEditPengajuan)
	router.Post("/status/:id", h.UpdateStatus)

}
