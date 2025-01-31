package services

import (
	"github.com/renatorrocha/auth-jwt-go/src/interfaces"
	"github.com/renatorrocha/auth-jwt-go/src/repositories"
	"github.com/renatorrocha/auth-jwt-go/src/services/user"
)

type ServiceContainer struct {
	UserService interfaces.UserService
}

func GetServices(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		UserService: user.NewUserService(repositoryContainer),
	}
}
