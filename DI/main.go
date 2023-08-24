package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/weslyramalho/GO/DI/product"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)

	usecase := product.NewProductUseCase(repository)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)
}
