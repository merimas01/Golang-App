package interfaces

import (
	"Golang-App/models"
)

type ICRUDService[T any, Tinsert any, Tupdate any, TSearch any] interface {
	Create(entity *Tinsert) (*T, error)
	GetByID(id uint) (*T, error)
	Update(entity *Tupdate, id uint) (*T, error)
	GetAll(search *TSearch) (models.PagedResult[T], error)
	Delete(id uint) error
}
