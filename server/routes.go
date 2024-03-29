package server

func SetupRoutes(s *Server) {
	s.App.Post("/api/quiz", s.QuizHandler.CreateQuiz)
	s.App.Get("/api/quiz/:id", s.QuizHandler.GetQuizByID)


	s.App.Post("/api/user", s.UserHandler.CreateUser)
}
