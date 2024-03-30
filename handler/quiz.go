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
	quiz.AllTimePlays = 0
	quiz.DailyPlays = 0
	quiz.WeeklyPlays = 0
	quiz.MonthlyPlays = 0
	quiz.Rating = 0
	quiz.AmountOfRatings = 0
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

func (q *QuizHandler) IncrementPlays(c fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := q.service.IncrementPlays(intID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := fiber.Map{
		"message": "Plays incremented successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (q *QuizHandler) GetQuizzes(c fiber.Ctx) error {
	quizzes, err := q.service.GetQuizzes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) RateQuiz(c fiber.Ctx) error {
	type Rating struct {
		Rating float32 `json:"rating"`
	}
	rating := Rating{}
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := json.Unmarshal(c.Body(), &rating); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if rating.Rating < 1 || rating.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "rating must be between 1 and 5",
		})
	}
	err = q.service.RateQuiz(intID, rating.Rating)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := fiber.Map{
		"message": "Quiz rated successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (q *QuizHandler) GetRating(c fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	rating, err := q.service.GetRating(intID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(rating)	
}