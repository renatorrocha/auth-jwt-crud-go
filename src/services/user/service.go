package user

import (
	"errors"

	"github.com/renatorrocha/auth-jwt-go/src/interfaces"
	"github.com/renatorrocha/auth-jwt-go/src/repositories"
	"github.com/renatorrocha/auth-jwt-go/src/structs"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(repositoryContainer repositories.RepositoryContainer) UserService {
	return UserService{
		UserRepository: repositoryContainer.UserRepository,
	}
}

func (us UserService) Create(user structs.User) (*structs.User, error) {

	if user.Age < 18 {
		return nil, errors.New("user must be at least 18 years old")
	}

	return us.UserRepository.Create(user)
}

func (us UserService) GetAll() ([]structs.User, error) {

	return us.UserRepository.GetAll()
}

func (us UserService) FindByID(id string) (*structs.User, error) {

	return us.UserRepository.FindByID(id)
}

func (us UserService) Update(user structs.User) (*structs.User, error) {

	return us.UserRepository.Update(user)
}

func (us UserService) Delete(id string) error {

	return us.UserRepository.Delete(id)
}
