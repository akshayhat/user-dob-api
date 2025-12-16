package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akshayhat/user-dob-api/config"
	"github.com/akshayhat/user-dob-api/internal/handler"
	"github.com/akshayhat/user-dob-api/internal/repository"
	"github.com/akshayhat/user-dob-api/internal/routes"
	"github.com/akshayhat/user-dob-api/internal/service"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// DB
	db := config.ConnectDB()
	defer db.Close()

	// Layers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	routes.RegisterUserRoutes(app, userHandler)

	app.Listen(":3000")
}



