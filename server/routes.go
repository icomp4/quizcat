package server

func SetupRoutes(s *Server) {
	// quiz handlers
	s.App.Post("/api/quiz", s.QuizHandler.CreateQuiz)
	s.App.Get("/api/quiz/:id", s.QuizHandler.GetQuizByID)
	s.App.Put("/api/quiz/:id/complete", s.QuizHandler.IncrementPlays)
	s.App.Get("/api/quizzes", s.QuizHandler.GetQuizzes)
	s.App.Put("/api/quiz/:id/rate", s.QuizHandler.RateQuiz)
	s.App.Get("/api/quiz/:id/rating", s.QuizHandler.GetRating)


	// user handlers
	s.App.Post("/api/signup", s.UserHandler.CreateUser)

}
