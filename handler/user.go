package handler

import (
	"encoding/json"
	"log/slog"
	"quizcat/model"
	"quizcat/service"
	"quizcat/util"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type UserHandler struct {
	service service.UserService
	store   *session.Store
	logger  slog.Logger
}

func NewUserHandler(service service.UserService, session *session.Store, logger slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		store:   session,
		logger:  logger,
	}
}
func (u *UserHandler) writeErrorWithLog(c fiber.Ctx, status int, message string) error {
	u.logger.Error(message)
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

func (u *UserHandler) writeError(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
func (u *UserHandler) CreateUser(c fiber.Ctx) error {
	var user model.User
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return u.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := util.ValidateUser(&user); err != nil {
		return u.writeError(c, fiber.StatusBadRequest, err.Error())
	}

	if err := u.service.CreateUser(&user); err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Signup successful!",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (u *UserHandler) Login(c fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var l login
	if err := json.Unmarshal(c.Body(), &l); err != nil {
		return u.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if l.Username == "" || l.Password == "" {
		return u.writeError(c, fiber.StatusBadRequest, "missing required fields")
	}
	id, err := u.service.Login(l.Username, l.Password)
	if err != nil {
		return u.writeErrorWithLog(c, fiber.StatusUnauthorized, err.Error())
	}
	sess.Set("userID", id)
	sess.Set("isAuth", true)
	if err := sess.Save(); err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Login successful!",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
func (u *UserHandler) Logout(c fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	sess.Destroy()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful!",
	})
}
func(u *UserHandler) GetUser(c fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return u.writeError(c, fiber.StatusUnauthorized, err.Error())
	}
	users, err := u.service.GetUserByID(userID)
	if err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
func (u *UserHandler) IsAuth(c fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		return u.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	isAuth, ok := sess.Get("isAuth").(bool)
	if !ok {
		return u.writeError(c, fiber.StatusUnauthorized, "unauthorized")
	}
	userID, ok := sess.Get("userID").(int)
	if !ok {
		return u.writeError(c, fiber.StatusUnauthorized, "unauthorized")
	}

	resp := fiber.Map{
		"isAuth": isAuth,
		"userID": userID,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
