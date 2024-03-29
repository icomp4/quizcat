package main

import (
	"fmt"
	"log"
	"quizcat/database"
	
)

func main() {
	db := database.New()
	queries := []string{`DROP TABLE IF EXISTS options`, `DROP TABLE IF EXISTS questions`, `DROP TABLE IF EXISTS quizzes`, `DROP TABLE IF EXISTS users`}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Successfully dropped all tables in the database")
}
