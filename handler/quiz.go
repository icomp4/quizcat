package handler

import (
	"encoding/json"
	"quizcat/model"
	"quizcat/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type QuizHandler struct {
	service service.QuizService
}

func NewQuizHandler(service service.QuizService) *QuizHandler {
	return &QuizHandler{
		service: service,
	}
}

func (q *QuizHandler) CreateQuiz(c fiber.Ctx) error {
	var quiz model.Quiz
	if err := json.Unmarshal(c.Body(), &quiz); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := q.service.CreateQuiz(&quiz); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := fiber.Map{
		"message": "Quiz created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}	


func (q *QuizHandler) GetQuizByID(c fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	quiz, err := q.service.GetQuizByID(intID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(quiz)
}		