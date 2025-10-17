package repository

//direct communication with the database

import (
	"Golang-App/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// all UserRepository functions

func (r *UserRepository) CreateUser(input *models.UserInsert) (*models.User, error) {
	user := &models.User{Name: input.Name, Email: input.Email, Age: input.Age}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (r *UserRepository) UpdateUser(input *models.UserUpdate, id uint) (*models.User, error) {
	var user models.User

	// Find the user first
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Update fields
	user.Name = input.Name
	user.Age = input.Age
	user.Email = input.Email

	// Save updated user
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) DeleteUser(id uint) error {
	// Find the user first
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return err // user not found
	}

	// Delete the user
	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
