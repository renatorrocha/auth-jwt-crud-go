package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/renatorrocha/auth-jwt-go/src/interfaces"
	"github.com/renatorrocha/auth-jwt-go/src/services"
	"github.com/renatorrocha/auth-jwt-go/src/structs"
)

type UserHandler struct {
	router      fiber.Router
	userService interfaces.UserService
}

func NewUserHandler(router fiber.Router, serviceContainer services.ServiceContainer) UserHandler {
	return UserHandler{
		router:      router,
		userService: serviceContainer.UserService,
	}
}

func (uh UserHandler) SetRoutes() {
	group := uh.router.Group("/users")

	group.Post("", uh.Create)
	group.Get("", uh.GetAll)
	group.Get("/:id", uh.FindById)
	group.Put("/:id", uh.Update)
	group.Delete("/:id", uh.Delete)
}

func (uh UserHandler) Create(c *fiber.Ctx) error {
	var body structs.User

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Response{Message: err.Error(), Code: http.StatusBadRequest})
	}

	user, err := uh.userService.Create(body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.Response{Message: err.Error(), Code: http.StatusInternalServerError})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func (uh UserHandler) GetAll(c *fiber.Ctx) error {
	var users []structs.User

	users, err := uh.userService.GetAll()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Response{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (uh UserHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uh.userService.FindByID(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Response{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.Status(http.StatusOK).JSON(user)

}

func (uh UserHandler) Update(c *fiber.Ctx) error {
	var body structs.User
	id := c.Params("id")

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Response{Message: err.Error(), Code: http.StatusBadRequest})
	}

	existingUser, err := uh.userService.FindByID(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Response{Message: err.Error(), Code: http.StatusBadRequest})
	}

	body.ID = existingUser.ID

	user, err := uh.userService.Update(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.Response{Message: err.Error(), Code: http.StatusInternalServerError})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (uh UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := uh.userService.Delete(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(structs.Response{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.Status(http.StatusOK).JSON(structs.Response{Message: "User deleted successfully", Code: http.StatusOK})
}
