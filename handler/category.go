package handler

import (
	"encoding/json"
	"quizcat/service"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type CategoryHandler struct {
	service service.CategoryService
	store   *session.Store
}

func NewCategoryHandler(service service.CategoryService, session *session.Store) *CategoryHandler {
	return &CategoryHandler{
		service: service,
		store:   session,
	}
}

func (ch *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	var body request
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ch.service.CreateCategory(body.Name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := fiber.Map{
		"message": "Category created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (ch *CategoryHandler) GetCategories(c fiber.Ctx) error {
	categories, err := ch.service.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(categories)
}

func (ch *CategoryHandler) AssignCategoryToQuiz(c fiber.Ctx) error {
	type request struct {
		QuizID     int `json:"quiz_id"`
		CategoryID int `json:"category_id"`
	}
	var body request
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ch.service.AssignCategoryToQuiz(body.QuizID, body.CategoryID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(quizzes)
}