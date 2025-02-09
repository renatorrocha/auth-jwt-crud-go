package interfaces

import "github.com/renatorrocha/auth-jwt-go/src/structs"

type UserService interface {
	Create(user structs.User) (*structs.User, error)
	GetAll() ([]structs.User, error)
	FindByID(id string) (*structs.User, error)
	Update(user structs.User) (*structs.User, error)
	Delete(id string) error
}

type UserRepository interface {
	Create(user structs.User) (*structs.User, error)
	GetAll() ([]structs.User, error)
	FindByID(id string) (*structs.User, error)
	Update(user structs.User) (*structs.User, error)
	Delete(id string) error
}
