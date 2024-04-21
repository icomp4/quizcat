package server

import (
	"quizcat/auth"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func SetupRoutes(s *Server) {
	authMiddleware := auth.ValidateUser(s.Session)
	limter := limiter.New(limiter.Config{
		Max:        1,
		Expiration: 24 * time.Hour,
		KeyGenerator: func(c fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c fiber.Ctx) error {
			return c.SendString("You have reached the limit of requests")
		},
	})
	// quiz handlers
	s.App.Post("/api/quiz", s.QuizHandler.CreateQuiz)                           // create quiz
	s.App.Get("/api/quiz/:id", s.QuizHandler.GetQuizByID)                       // get quiz by id
	s.App.Put("/api/quiz/:id/complete", s.QuizHandler.IncrementPlays)           // increment plays
	s.App.Put("/api/quiz/:id/rate", s.QuizHandler.RateQuiz, authMiddleware, limter)     // rate quiz
	s.App.Get("/api/quiz/:id/rating", s.QuizHandler.GetRating)                  // get rating
	s.App.Get("/api/quizzes/top/:period", s.QuizHandler.GetTopQuizzesPerPeriod) // get top quizzes per period
	s.App.Get("/api/quizzes/search", s.QuizHandler.SearchQuizzes)               // search quizzes
	s.App.Delete("/api/quiz/:id", s.QuizHandler.DeleteQuiz, authMiddleware)     // delete quiz

	// user handlers
	s.App.Post("/api/signup", s.UserHandler.CreateUser) // create user
	s.App.Post("/api/login", s.UserHandler.Login)       // login
	s.App.Get("/api/isAuth", s.UserHandler.IsAuth)      // check if user is authenticated

	// category handlers
	s.App.Post("/api/category", s.CategoryHandler.CreateCategory, authMiddleware)                        // create category
	s.App.Get("/api/categories", s.CategoryHandler.GetCategories)                                        // get categories
	s.App.Post("/api/quiz/:quizID/category/:id", s.CategoryHandler.AssignCategoryToQuiz, authMiddleware) // assign category to quiz
	s.App.Get("/api/quizzes/:category", s.CategoryHandler.GetQuizzesByCategory)                          // get quizzes by category

	// question handlers
	s.App.Post("/api/quiz/:id/question", s.QuestionHandler.CreateQuestion, authMiddleware) // create question
	s.App.Put("/api/question/:id", s.QuestionHandler.EditQuestion, authMiddleware)         // edit question
	s.App.Delete("/api/question/:id", s.QuestionHandler.DeleteQuestion, authMiddleware)    // delete question

	// option handlers
	s.App.Post("/api/question/:id/option", s.QuestionHandler.CreateOption, authMiddleware)    // create option
	s.App.Put("/api/option/:id/text", s.QuestionHandler.EditOption, authMiddleware)           // edit option
	s.App.Put("/api/option/:id/correct", s.QuestionHandler.EditOptionCorrect, authMiddleware) // edit option correct
	s.App.Delete("/api/option/:id", s.QuestionHandler.DeleteOption, authMiddleware)           // delete option

}
