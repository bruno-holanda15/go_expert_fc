package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/dto"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/infra/database"
)

type ProductHandler struct {
	// usando interface para não injetar diretamente pacote ??
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// deveria estar dentro da camada application, como no usecase
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// deveria estar dentro da camada application, como no usecase
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
