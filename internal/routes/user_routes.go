package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akshayhat/user-dob-api/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, handler *handler.UserHandler) {
	users := app.Group("/users")

	users.Post("", handler.CreateUser)
	users.Get("", handler.ListUsers)
	users.Get("/:id", handler.GetUserByID)
}
