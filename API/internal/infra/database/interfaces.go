package database

import "github.com/weslyramalho/GO/tree/main/API/internal/entity"

type UserInterface interface {
	Create(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(Product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(Product *entity.Product) error
	Delete(id string) error
}
