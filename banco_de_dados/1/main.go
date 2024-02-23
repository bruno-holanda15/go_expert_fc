package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_expert_br")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	prod := NewProduct("Emulador", 400)
	
	err = insertProduct(db, prod)
	if err != nil {
		panic(err)
	}

	prod.Price = 341
	err = updateProduct(db, prod)
	if err != nil {
		panic(err)
	}

	// p, err := selectProduct(db, prod.ID)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("O Produto com nome %s -> Custa R$%v", p.Name, p.Price)
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("O Produto com nome %s -> Custa R$%v\n", p.Name, p.Price)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	// Adicionamos o exemple de Exec com context para controlarmos o timeout da execução como exemplo
	// _, err = stmt.ExecContext(ctx, product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price= ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	// Abaixo podemos utilizar o context dentro da consulta para limitar o tempo como no desafio
	// err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		panic(err)
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}