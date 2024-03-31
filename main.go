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
	"github.com/gofiber/fiber/v3/middleware/session"
	_ "github.com/joho/godotenv/autoload"
)
func main() {
	db := database.New()
	store := session.New()
	quizService := service.NewQuizService(db)
	userService := service.NewUserService(db)
	categoryService := service.NewCategoryService(db)

	quizHandler := handler.NewQuizHandler(*quizService)
	userHandler := handler.NewUserHandler(*userService, store)
	categoryHandler := handler.NewCategoryHandler(*categoryService, store)

	app := server.New(fiber.New(), quizHandler, userHandler, categoryHandler, store)
	app.App.Use(logger.New())
	server.SetupRoutes(app)
	log.Fatal(app.App.Listen(":" + os.Getenv("PORT")))
	util.StartPlayCountResetScheduler(db)
	select {}
}
