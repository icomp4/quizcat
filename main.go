package main

import (
	"log"
	"os"
	"quizcat/database"
	"quizcat/handler"
	"quizcat/server"
	"quizcat/service"
	"quizcat/util"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	quizService := service.NewQuizService(db)
	userService := service.NewUserService(db)

	quizHandler := handler.NewQuizHandler(*quizService)
	userHandler := handler.NewUserHandler(*userService)

	app := server.New(fiber.New(), *quizHandler, *userHandler)
	app.App.Use(logger.New())
	server.SetupRoutes(app)
	log.Fatal(app.App.Listen(":" + os.Getenv("PORT")))
	util.StartPlayCountResetScheduler(db)
	select {}
}
