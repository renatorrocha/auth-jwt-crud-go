package structs

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Age   uint8  `json:"age"`
}
