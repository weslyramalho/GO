package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weslyramalho/GO/tree/main/API/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	porductDB := NewProduct(db)
	err = porductDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), 12.1)
		assert.NoError(t, err)
		db.Create(product)
	}
	porductDB := NewProduct(db)
	products, err := porductDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	//assert.Equal(t, "Product 1", products[0].Name)
	//assert.Equal(t, "Product 10", products[1].Name)

}
