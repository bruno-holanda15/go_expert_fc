package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("PC", 1200.0)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.NotEmpty(t, productFound.ID)
}

func TestFinalAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProductDB(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Teclado", 550.0)
	assert.NoError(t, err)
	assert.NotEmpty(t, product)

	productDB := NewProductDB(db)
	db.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)

	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

	_, err = productDB.FindById("fake-uid")
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("IPhone", 13000.0)
	assert.NoError(t, err)

	db.Create(product)
	product.Price = 14000.0

	productDB := NewProductDB(db)
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Price, productFound.Price)

	productNotInsertedInDB, _ := entity.NewProduct("Moto", 13000.0)

	err = productDB.Update(productNotInsertedInDB)
	assert.Error(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProductDB(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	err = productDB.Delete("fake-uuid")
	assert.Error(t, err)
}
