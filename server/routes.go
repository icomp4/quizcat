package server

import "quizcat/auth"

func SetupRoutes(s *Server) {
	authMiddleware := auth.ValidateUser(s.Session)

	// quiz handlers
	// quizGroup := s.App.Group("/api/quiz", authMiddleware)
	// quizGroup.Post("/", s.QuizHandler.CreateQuiz)
	// quizGroup.Get("/:id", s.QuizHandler.GetQuizByID)
	// quizGroup.Put("/:id/complete", s.QuizHandler.IncrementPlays)
	// quizGroup.Get("/quizzes", s.QuizHandler.GetQuizzes)
	// quizGroup.Put("/:id/rate", s.QuizHandler.RateQuiz)
	// quizGroup.Get("/:id/rating", s.QuizHandler.GetRating)

	s.App.Post("/api/quiz", s.QuizHandler.CreateQuiz)
	s.App.Get("/api/quiz/:id", s.QuizHandler.GetQuizByID)
	s.App.Put("/api/quiz/:id/complete", s.QuizHandler.IncrementPlays)
	s.App.Put("/api/quiz/:id/rate", s.QuizHandler.RateQuiz)
	s.App.Get("/api/quiz/:id/rating", s.QuizHandler.GetRating)
	s.App.Get("/api/quizzes/top/:period", s.QuizHandler.GetTopQuizzesPerPeriod)
	s.App.Get("/api/quizzes/search", s.QuizHandler.SearchQuizzes)

	// user handlers
	s.App.Post("/api/signup", s.UserHandler.CreateUser)
	s.App.Post("/api/login", s.UserHandler.Login)
	s.App.Get("/api/authTest", s.UserHandler.Test, authMiddleware)

	// category handlers
	s.App.Post("/api/category", s.CategoryHandler.CreateCategory)
	s.App.Get("/api/categories", s.CategoryHandler.GetCategories)
	s.App.Post("/api/quiz/:id/category", s.CategoryHandler.AssignCategoryToQuiz)
	s.App.Get("/api/quizzes/category/:category", s.CategoryHandler.GetQuizzesByCategory)

}
