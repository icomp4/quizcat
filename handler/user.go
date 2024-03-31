package handler

import (
	"encoding/json"
	"quizcat/model"
	"quizcat/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type UserHandler struct {
	service service.UserService
	store *session.Store
}

func NewUserHandler(service service.UserService, session *session.Store) *UserHandler {
    return &UserHandler{
        service: service,
        store: session,
    }
}

func (u *UserHandler) CreateUser(c fiber.Ctx) error {
	var user model.User
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if user.Email == "" || user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "missing required fields",
		})
	}

	if err := u.service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := fiber.Map{
		"message": "Signup successful!",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (u *UserHandler) Login(c fiber.Ctx) error {
	sess, err := u.store.Get(c)
    if err != nil {
        panic(err)
    }
	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var l login
	if err := json.Unmarshal(c.Body(), &l); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if l.Username == "" || l.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "missing required fields",
		})
	}
	id, err := u.service.Login(l.Username, l.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	sess.Set("userID", id)
	sess.Set("isAuth", true)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := fiber.Map{
		"message": "Login successful!",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (u *UserHandler) Test(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}