//go:build wireinject
// +build wireiject

package main

import (
	"database/sql"
	"go_di_wire/product"

	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}