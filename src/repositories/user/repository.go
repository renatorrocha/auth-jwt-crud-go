package user

import (
	"github.com/renatorrocha/auth-jwt-go/src/structs"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur UserRepository) Create(user structs.User) (*structs.User, error) {
	err := ur.db.Create(&user).Error
	return &user, err
}

func (ur UserRepository) GetAll() ([]structs.User, error) {
	var users []structs.User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur UserRepository) FindByID(id string) (*structs.User, error) {
	var user structs.User
	err := ur.db.First(&user, id).Error
	return &user, err
}

func (ur UserRepository) Update(user structs.User) (*structs.User, error) {
	err := ur.db.Save(&user).Error
	return &user, err
}

func (ur UserRepository) Delete(id string) error {
	err := ur.db.Delete(&structs.User{}, id).Error

	return err
}
