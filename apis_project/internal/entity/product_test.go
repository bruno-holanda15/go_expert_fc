package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("ps5", 4000)

	assert.NotEmpty(t, product.Name)
	assert.NotNil(t, product.Price)
	assert.Nil(t, err)
	assert.Equal(t, "ps5", product.Name)
	assert.Equal(t, 4000, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 1000)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("PC", 0)

	assert.Nil(t, product)
	assert.Equal(t,  ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Nintendo", 2000)

	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.NoError(t, product.Validate())
}