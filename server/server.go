package server

import (
	"quizcat/handler"

	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App *fiber.App
	Session *session.Store
	QuizHandler handler.QuizHandler
	UserHandler handler.UserHandler
	QuestionHandler handler.QuestionHandler
	CategoryHandler handler.CategoryHandler
}

func New(app *fiber.App, quizHandler *handler.QuizHandler, userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler,
	questionHandler *handler.QuestionHandler, session *session.Store) *Server {
	return &Server{
		App:         app,
		QuizHandler: *quizHandler,
		UserHandler: *userHandler,
		CategoryHandler: *categoryHandler,
		QuestionHandler: *questionHandler,
		Session: session,
	}
}
