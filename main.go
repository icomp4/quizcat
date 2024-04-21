package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"quizcat/database"
	"quizcat/handler"
	"quizcat/server"
	"quizcat/service"
	"quizcat/util"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/session"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	store := session.New(session.Config{
		CookieSecure:   true,
		CookieHTTPOnly: true,
		Expiration:     30 * (time.Hour * 24),
	})
	logger := slog.New(slog.Default().Handler())
	quizService := service.NewQuizService(db)
	userService := service.NewUserService(db)
	categoryService := service.NewCategoryService(db)
	questionService := service.NewQuestionService(db)

	quizHandler := handler.NewQuizHandler(*quizService, store, *logger)
	userHandler := handler.NewUserHandler(*userService, store, *logger)
	categoryHandler := handler.NewCategoryHandler(*categoryService, store, *logger)
	questionHandler := handler.NewQuestionHandler(*questionService, store, *logger)

	app := server.New(fiber.New(), quizHandler, userHandler, categoryHandler, questionHandler, store)
	app.App.Use(cors.New(cors.Config{
		AllowOrigins:     "https://quizcat-production.up.railway.app",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
	}))
	app.App.Static("/", "./frontend/dist")
	server.SetupRoutes(app)
	port := os.Getenv("PORT")
	portStr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Fatal(app.App.Listen(portStr))
	util.StartPlayCountResetScheduler(db)
	select {}
}
