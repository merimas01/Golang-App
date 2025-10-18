package interfaces

import (
	"Golang-App/models"
	searchobjects "Golang-App/models/search_objects"
)

type IUserService interface {
	ICRUDService[models.User, models.UserInsert, models.UserUpdate, searchobjects.BaseSearchObject] //inherits this interface
	//new methods for user...
}
