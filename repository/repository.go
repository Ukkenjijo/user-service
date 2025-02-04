package repository

import (
	"github.com/Ukkenjijo/user-service/service"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user service.User) (service.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user service.User) (service.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return service.User{}, err
	}
	return user, nil
}
