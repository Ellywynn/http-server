package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
)

type UserService struct {
	repository models.UserRepository
}

func NewUserService(repo *models.UserRepository) models.UserService {
	return &UserService{
		repository: *repo,
	}
}

func (us *UserService) SignUp(user *models.User) (int, error) {
	return us.repository.Create(user)
}

func (us *UserService) GetByEmail(email string) (*models.User, error) {
	return us.repository.FindByEmail(email)
}

func (us *UserService) GetByUsername(username string) (*models.User, error) {
	return us.repository.FindByUsername(username)
}

func (us *UserService) GetById(userId int) (*models.User, error) {
	return us.repository.FindById(userId)
}

func (us *UserService) GetAll() (*[]models.User, error) {
	return us.repository.GetAll()
}

func (us *UserService) Update(userId int) error {
	return us.repository.Update(userId)
}

func (us *UserService) Delete(userId int) error {
	return us.repository.Delete(userId)
}
