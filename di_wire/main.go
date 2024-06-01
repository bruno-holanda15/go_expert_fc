package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// repository := product.NewProductRepository(db)
	// usecase := product.NewProductUseCase(repository)

	// utilizando wire
	usecase := NewProductUseCase(db)

	p, err := usecase.GetProductById(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name)
}
