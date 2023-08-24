//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/weslyramalho/GO/DI/product"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
