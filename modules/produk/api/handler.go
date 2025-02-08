// handler.go

package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/donnyirianto/be-stock-app/modules/produk/usecase"
)

type ProdukHandler struct {
	produk   usecase.ProdukUsecase
	validate *validator.Validate
}

func NewProdukHandler(produk usecase.ProdukUsecase) *ProdukHandler {
	return &ProdukHandler{
		produk:   produk,
		validate: validator.New(),
	}
}

func (h *ProdukHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetProduk)
	router.Post("/", h.AddProduk)
	router.Get("/edit/:id", h.EditProduk)
	router.Post("/edit/:id", h.SaveEditProduk)
	router.Delete("/:id", h.DeleteProduk)
}
