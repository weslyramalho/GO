package database

import "github.com/weslyramalho/GO/tree/main/API/internal/entity"

type UserInterface interface {
	Create(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
