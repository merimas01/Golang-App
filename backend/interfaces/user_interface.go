package interfaces

import "Golang-App/models"

type IUserService interface {
	ICRUDService[models.User, models.UserInsert, models.UserUpdate] //inherits this interface
	//new methods for user...
}
