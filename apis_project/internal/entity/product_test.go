package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("ps5", 400.0)

	assert.NotEmpty(t, product.Name)
	assert.NotNil(t, product.Price)
	assert.Nil(t, err)
	assert.Equal(t, "ps5", product.Name)
	assert.Equal(t, 400.0, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100.0)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("PC", 0.0)

	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Nintendo", 200.0)

	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.NoError(t, product.Validate())
}

func TestProductBatch(t *testing.T) {
	type dataProduct struct {
		name        string
		price       float64
		errExpected error
	}
	products := []dataProduct{
		{"Nintendo", 0.0, errors.New("invalid price")},
		{"PC", -120.0, errors.New("invalid price")},
		{"", 430.0, errors.New("name is required")},
	}

	for _, product := range products {
		_, err := NewProduct(product.name, product.price)

		if err != nil && err.Error() != product.errExpected.Error() {
			t.Errorf("Expected error '%s', but got '%s'", product.errExpected, err)
		}
	}
}
