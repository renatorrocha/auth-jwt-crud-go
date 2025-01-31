package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatorrocha/auth-jwt-go/src/handlers/user"
	"github.com/renatorrocha/auth-jwt-go/src/services"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	user.NewUserHandler(router, serviceContainer).SetRoutes()
}
