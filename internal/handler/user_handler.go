package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/akshayhat/user-dob-api/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// -------- Request DTO --------
type CreateUserRequest struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

// -------- Handlers --------

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.service.CreateUser(c.Context(), req.Name, req.Dob)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GET /users/:id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	user, age, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  age,
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(users)
}
