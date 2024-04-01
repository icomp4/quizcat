package handler

import (
	"encoding/json"
	"log/slog"
	"quizcat/model"
	"quizcat/service"
	"quizcat/util"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type QuestionHandler struct {
	service service.QuestionService
	store   *session.Store
	logger  slog.Logger
}

func NewQuestionHandler(service service.QuestionService, session *session.Store, logger slog.Logger) *QuestionHandler {
	return &QuestionHandler{
		service: service,
		store:   session,
		logger:  logger,
	}
}
func (q *QuestionHandler) writeErrorWithLog(c fiber.Ctx, status int, message string) error {
	q.logger.Error(message)
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
func (q *QuestionHandler) writeError(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
func (q *QuestionHandler) CreateQuestion(c fiber.Ctx) error {
	quizID, err := strconv.Atoi(c.Params("quizID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess,err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	var question model.Question
	if err := json.Unmarshal(c.Body(), &question); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.CreateQuestion(userID, quizID, question.Text); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Question created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
func (q *QuestionHandler) EditQuestion(c fiber.Ctx) error {
	questionID, err := strconv.Atoi(c.Params("questionID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	var question model.Question
	if err := json.Unmarshal(c.Body(), &question); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.EditQuestionText(userID, questionID, question.Text); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Question edited successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
func (q *QuestionHandler) DeleteQuestion(c fiber.Ctx) error {
	questionID, err := strconv.Atoi(c.Params("questionID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	if err := q.service.DeleteQuestion(userID, questionID); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Question deleted successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
func(q *QuestionHandler) CreateOption(c fiber.Ctx) error {
	questionID, err := strconv.Atoi(c.Params("questionID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	var option model.Option
	if err := json.Unmarshal(c.Body(), &option); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.CreateOption(userID, questionID, option.Text, option.IsCorrect); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Option created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
func (q *QuestionHandler) EditOption(c fiber.Ctx) error {
	optionID, err := strconv.Atoi(c.Params("optionID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	var option model.Option
	if err := json.Unmarshal(c.Body(), &option); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.EditOptionText(userID, optionID, option.Text); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Option edited successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
func (q *QuestionHandler) EditOptionCorrect(c fiber.Ctx) error {
	optionID, err := strconv.Atoi(c.Params("optionID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	var option model.Option
	if err := json.Unmarshal(c.Body(), &option); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	if err := q.service.EditOptionCorrect(userID, optionID, option.IsCorrect); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Option edited successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
func (q *QuestionHandler) DeleteOption(c fiber.Ctx) error {
	optionID, err := strconv.Atoi(c.Params("optionID"))
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	sess, err := q.store.Get(c)
	if err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	userID, err := util.GetUserIDFromSession(*sess)
	if err != nil {
		return q.writeError(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	if err := q.service.DeleteOption(userID, optionID); err != nil {
		return q.writeErrorWithLog(c, fiber.StatusInternalServerError, err.Error())
	}
	resp := fiber.Map{
		"message": "Option deleted successfully",
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
