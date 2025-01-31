package database

import (
	"fmt"
	"strconv"

	"github.com/renatorrocha/auth-jwt-go/src/config"
	"github.com/renatorrocha/auth-jwt-go/src/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	p := config.Config("DATABASE_PORT")

	// Convert port to uint (10 is the base, 32 is the bit size)
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse the database port")
	}

	dsn := fmt.Sprintf("host=db port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Config("DATABASE_USER"),
		config.Config("DATABASE_PASSWORD"),
		config.Config("DATABASE_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected successfully")
	DB.AutoMigrate(&structs.User{})
	fmt.Println("Database migrated successfully")
}
