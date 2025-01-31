package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/renatorrocha/auth-jwt-go/src/database"
	"github.com/renatorrocha/auth-jwt-go/src/handlers"
	"github.com/renatorrocha/auth-jwt-go/src/repositories"
	"github.com/renatorrocha/auth-jwt-go/src/services"
)

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	repositoryContainer := repositories.GetRepositories(database.DB)
	servicesContainer := services.GetServices(repositoryContainer)
	handlers.NewHandlerContainer(app, servicesContainer)

	log.Fatal(app.Listen(":3000"))
}
