package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProductById(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProductById(id int) (Product, error) {
	return Product{
		ID: id, 
		Name: "Product Name test",
	},	 nil
}

type ProductRepositoryTxt struct {
	db *sql.DB
}

func NewProductRepositoryTxt(db *sql.DB) *ProductRepositoryTxt {
	return &ProductRepositoryTxt{
		db: db,
	}
}

func (r *ProductRepositoryTxt) GetProductById(id int) (Product, error) {
	return Product{
		ID: id, 
		Name: "Product Name test",
	},	 nil
}
