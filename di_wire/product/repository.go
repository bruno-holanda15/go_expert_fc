package product

import "database/sql"

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
