package handler

import (
	"encoding/json"
	"log/slog"
	"quizcat/service"
	"quizcat/util"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type CategoryHandler struct {
	service service.CategoryService
	store   *session.Store
	logger slog.Logger
}

func NewCategoryHandler(service service.CategoryService, session *session.Store, logger slog.Logger) *CategoryHandler {
	return &CategoryHandler{
		service: service,
		store:   session,
		logger: logger,
	}
}
func(ch *CategoryHandler) writeErrorWithLog(c fiber.Ctx, status int, message string) error {
    ch.logger.Error(message)
    return c.Status(status).JSON(fiber.Map{
        "error": message,
    })
}

func (ch *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	var body request
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := ch.service.CreateCategory(body.Name); err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Category created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (ch *CategoryHandler) GetCategories(c fiber.Ctx) error {
	categories, err := ch.service.GetCategories()
	if err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(categories)
}

func (ch *CategoryHandler) AssignCategoryToQuiz(c fiber.Ctx) error {
	category := c.Params("id")
	quiz := c.Params("quizID")
	categoryID, err := strconv.Atoi(category)
	if err != nil{
		return ch.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	quizID, err := strconv.Atoi(quiz)
	if err != nil{
		return ch.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}

	sess, err := ch.store.Get(c)
	if err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}

	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusUnauthorized, err.Error())
	}
	if err := ch.service.AssignCategoryToQuiz(userID, quizID, categoryID); err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Category assigned to quiz successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (ch *CategoryHandler) GetQuizzesByCategory(c fiber.Ctx) error {
	category := c.Params("category")
	category = strings.Replace(category, "%20", " ", -1)
	quizzes, err := ch.service.GetQuizzesByCategory(category)
	if err != nil {
		return ch.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(quizzes)
}