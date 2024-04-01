package handler

import (
	"encoding/json"
	"log/slog"
	"quizcat/model"
	"quizcat/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type QuizHandler struct {
	service service.QuizService
	store   *session.Store
	logger  slog.Logger
}

func NewQuizHandler(service service.QuizService,session *session.Store, logger slog.Logger) *QuizHandler {
	return &QuizHandler{
		service: service,
		store:   session,
		logger:  logger,
	}
}
func (q *QuizHandler) writeErrorWithLog(c fiber.Ctx, status int, message string) error {
	q.logger.Error(message)
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

func (q *QuizHandler) writeError(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
func (q *QuizHandler) CreateQuiz(c fiber.Ctx) error {
	var quiz model.Quiz
	if err := json.Unmarshal(c.Body(), &quiz); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	quiz.AllTimePlays = 0
	quiz.DailyPlays = 0
	quiz.WeeklyPlays = 0
	quiz.MonthlyPlays = 0
	quiz.Rating = 0
	quiz.AmountOfRatings = 0
	if err := q.service.CreateQuiz(&quiz); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
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
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	quiz, err := q.service.GetQuizByID(intID)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(quiz)
}

func (q *QuizHandler) IncrementPlays(c fiber.Ctx) error {
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.IncrementPlays(intID); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Plays incremented successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
func (q *QuizHandler) RateQuiz(c fiber.Ctx) error {
	type Rating struct {
		Rating float32 `json:"rating"`
	}
	rating := Rating{}
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := json.Unmarshal(c.Body(), &rating); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if rating.Rating < 1 || rating.Rating > 5 {
		return q.writeError(c, fiber.StatusBadRequest, "Rating must be between 1 and 5")
	}
	err = q.service.RateQuiz(intID, rating.Rating)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Quiz rated successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (q *QuizHandler) GetRating(c fiber.Ctx) error {
	type rating struct {
		Rating float32 `json:"rating"`
	}
	var ratingResp rating
	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	ratingint, err := q.service.GetRating(intID)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	ratingResp.Rating = ratingint
	return c.JSON(ratingResp)
}

func (q *QuizHandler) GetTopQuizzesPerPeriod(c fiber.Ctx) error {
	period := c.Params("period")
	quizzes, err := q.service.GetTopQuizzesPerPeriod(period)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) SearchQuizzes(c fiber.Ctx) error {
	search := c.Query("param")
	quizzes, err := q.service.SearchQuizzes(search)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) DeleteQuiz(c fiber.Ctx) error {
	id := c.Params("id")
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, ok := sess.Get("userID").(int)
	if !ok {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.DeleteQuiz(userID, intID); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Quiz deleted successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
