package main

import (
	"log"
	"log/slog"
	"os"
	"quizcat/database"
	"quizcat/handler"
	"quizcat/server"
	"quizcat/service"
	"quizcat/util"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	store := session.New()
	logger := slog.New(slog.Default().Handler())
	quizService := service.NewQuizService(db)
	userService := service.NewUserService(db)
	categoryService := service.NewCategoryService(db)

	quizHandler := handler.NewQuizHandler(*quizService, *logger)
	userHandler := handler.NewUserHandler(*userService, store, *logger)
	categoryHandler := handler.NewCategoryHandler(*categoryService, store, *logger)

	app := server.New(fiber.New(), quizHandler, userHandler, categoryHandler, store)
	server.SetupRoutes(app)
	log.Fatal(app.App.Listen(":" + os.Getenv("PORT")))
	util.StartPlayCountResetScheduler(db)
	select {}
}
