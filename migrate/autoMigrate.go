package migrate

import (
	"database/sql"
)

func AutoMigrate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS quizzes (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			author_id INTEGER REFERENCES users(id),
			rating FLOAT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS questions (
			id SERIAL PRIMARY KEY,
			quiz_id INTEGER REFERENCES quizzes(id),
			text TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS options (
			id SERIAL PRIMARY KEY,
			question_id INTEGER REFERENCES questions(id),
			text TEXT NOT NULL,
			is_correct BOOLEAN NOT NULL
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
