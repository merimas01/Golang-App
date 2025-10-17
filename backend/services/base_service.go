package services

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BaseService[T any, Tinsert any, Tupdate any] struct {
	DB *gorm.DB
}

// Create
func (r *BaseService[T, Tinsert, Tupdate]) Create(input *Tinsert) (*T, error) {
	var entity T

	// Copy matching fields from Tinsert to T
	if err := copier.Copy(&entity, input); err != nil {
		return nil, err
	}

	// Save entity to DB
	if err := r.DB.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Update
func (r *BaseService[T, Tinsert, Tupdate]) Update(input *Tupdate, id uint) (*T, error) {
	var entity T

	// Find the record first
	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}

	// Update only non-zero fields from input
	if err := r.DB.Model(&entity).Updates(input).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// GetByID
func (r *BaseService[T, Tinsert, Tupdate]) GetByID(id uint) (*T, error) {
	var entity T
	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetAll
func (r *BaseService[T, Tinsert, Tupdate]) GetAll() ([]T, error) {
	var entities []T
	if err := r.DB.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Delete
func (r *BaseService[T, Tinsert, Tupdate]) Delete(id uint) error {
	var entity T

	// Check if the record exists
	if err := r.DB.First(&entity, id).Error; err != nil {
		return err // not found
	}

	// Delete the record
	if err := r.DB.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}
