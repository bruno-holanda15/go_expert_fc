package entity

import (
	"errors"
	"time"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	CreateAT time.Time    `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	if name == "" {
		return nil, ErrNameIsRequired
	}

	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	product := &Product{
		ID:       entity.NewID(),
		Name:     name,
		Price:    price,
		CreateAT: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
