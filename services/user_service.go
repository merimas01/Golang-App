package services

import (
	"Golang-App/models"
	"Golang-App/repository"
)

type UserService struct {
	Repo *repository.UserRepository //package repository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.UserInsert) (*models.User, error) {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.UserUpdate, id uint) (*models.User, error) {
	return s.Repo.UpdateUser(user, id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}
