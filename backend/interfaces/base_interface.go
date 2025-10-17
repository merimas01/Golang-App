package interfaces

type ICRUDService[T any, Tinsert any, Tupdate any] interface {
	Create(entity *Tinsert) (*T, error)
	GetByID(id uint) (*T, error)
	Update(entity *Tupdate, id uint) (*T, error)
	GetAll() ([]T, error)
	Delete(id uint) error
}
