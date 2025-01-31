package repositories

import (
	"github.com/renatorrocha/auth-jwt-go/src/interfaces"
	"github.com/renatorrocha/auth-jwt-go/src/repositories/user"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	UserRepository interfaces.UserRepository
}

func GetRepositories(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		UserRepository: user.NewUserRepository(db),
	}
}
