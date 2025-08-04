package repository

import (
    "github.com/Anujc07/real-estate/internal/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
    return r.DB.Create(user).Error
}
	
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    result := r.DB.Where("email = ?", email).First(&user)
    return &user, result.Error
}
