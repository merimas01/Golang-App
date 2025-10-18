package services

import (
	"Golang-App/interfaces"
	"Golang-App/models"
	searchobjects "Golang-App/models/search_objects"
)

type UserService struct {
	Repo *BaseService[models.User, models.UserInsert, models.UserUpdate, searchobjects.BaseSearchObject]
}

// UserService implements IUserService
var _ interfaces.IUserService = &UserService{}

// Constructor
func NewUserService(repo *BaseService[models.User, models.UserInsert, models.UserUpdate, searchobjects.BaseSearchObject]) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Create(input *models.UserInsert) (*models.User, error) {
	return s.Repo.Create(input)
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) Update(input *models.UserUpdate, id uint) (*models.User, error) {
	return s.Repo.Update(input, id)
}

func (s *UserService) GetAll(search *searchobjects.BaseSearchObject) (models.PagedResult[models.User], error) {
	return s.Repo.GetAll(search)
}

func (s *UserService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

//user-specific methods
