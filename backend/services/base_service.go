package services

import (
	"Golang-App/models"
	"reflect"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BaseService[T any, Tinsert any, Tupdate any, Tsearch any] struct {
	DB *gorm.DB
}

// Create
func (r *BaseService[T, Tinsert, Tupdate, Tsearch]) Create(input *Tinsert) (*T, error) {
	var entity T

	if err := copier.Copy(&entity, input); err != nil {
		return nil, err
	}

	if err := r.DB.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Update
func (r *BaseService[T, Tinsert, Tupdate, Tsearch]) Update(input *Tupdate, id uint) (*T, error) {
	var entity T

	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Model(&entity).Updates(input).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// GetByID
func (r *BaseService[T, Tinsert, Tupdate, Tsearch]) GetByID(id uint) (*T, error) {
	var entity T
	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetAll
func (r *BaseService[T, Tinsert, Tupdate, Tsearch]) GetAll(search *Tsearch) (models.PagedResult[T], error) {
	var entities []T
	var count int64

	query := r.DB.Model(new(T)) //Model govori sa kojom DB tabelom cemo raditi

	if err := query.Count(&count).Error; err != nil {
		return models.PagedResult[T]{}, err
	}

	var page int
	pageSize := 10
	if v := reflect.ValueOf(search); v.IsValid() && !v.IsZero() {
		if v.Kind() == reflect.Ptr {
			v = v.Elem() //dereferencira search obj
		}

		pageField := v.FieldByName("Page")
		pageSizeField := v.FieldByName("PageSize")

		if pageField.IsValid() && !pageField.IsZero() {
			page = int(pageField.Int())
		}
		if pageSizeField.IsValid() && !pageSizeField.IsZero() {
			pageSize = int(pageSizeField.Int())
		}
	}

	if pageSize > 0 {
		query = query.Limit(pageSize).Offset(page * pageSize)
	}

	if err := query.Find(&entities).Error; err != nil {
		return models.PagedResult[T]{}, err
	}

	return models.PagedResult[T]{
		Result: entities,
		Count:  int(count),
	}, nil

}

// Delete
func (r *BaseService[T, Tinsert, Tupdate, Tsearch]) Delete(id uint) error {
	var entity T

	if err := r.DB.First(&entity, id).Error; err != nil {
		return err // not found
	}

	if err := r.DB.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}
