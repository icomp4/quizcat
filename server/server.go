package server

import (
	"quizcat/handler"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App         *fiber.App
	QuizHandler handler.QuizHandler
	UserHandler handler.UserHandler
}

func New(app *fiber.App, quizHandler handler.QuizHandler, userHandler handler.UserHandler) *Server {
	return &Server{
		App:         app,
		QuizHandler: quizHandler,
		UserHandler: userHandler,
	}
}
