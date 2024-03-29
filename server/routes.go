package server

func SetupRoutes(s *Server) {
	s.App.Post("/api/quiz", s.QuizHandler.CreateQuiz)
	s.App.Get("/api/quiz/:id", s.QuizHandler.GetQuizByID)
	s.App.Put("/api/quiz/:id/complete", s.QuizHandler.IncrementPlays)


	s.App.Post("/api/signup", s.UserHandler.CreateUser)

}
