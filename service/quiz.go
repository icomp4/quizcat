package service

import (
	"database/sql"
	"fmt"
	"quizcat/model"
)

type QuizService struct {
	db *sql.DB
}

func NewQuizService(db *sql.DB) *QuizService {
	return &QuizService{
		db: db}
}

func (q *QuizService) CreateQuiz(quiz *model.Quiz) error {
	tx, err := q.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var quizID uint
	err = tx.QueryRow("INSERT INTO quizzes (title, author_id, rating) VALUES ($1, $2, $3) RETURNING id", quiz.Title, quiz.AuthorID, quiz.Rating).Scan(&quizID)
	if err != nil {
		return fmt.Errorf("error creating quiz: %w", err)
	}

	for _, question := range quiz.Questions {
		var questionID uint
		err = tx.QueryRow("INSERT INTO questions (quiz_id, text) VALUES ($1, $2) RETURNING id", quizID, question.Text).Scan(&questionID)
		if err != nil {
			return fmt.Errorf("error creating question: %w", err)
		}
		for _, option := range question.Options {
			if _, err := tx.Exec("INSERT INTO options (question_id, text, is_correct) VALUES ($1, $2, $3)", questionID, option.Text, option.IsCorrect); err != nil {
				return fmt.Errorf("error creating option: %w", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (q *QuizService) GetQuizByID(id int) (*model.Quiz, error) {
	quiz := &model.Quiz{}
	err := q.db.QueryRow("SELECT id, title, author_id, rating FROM quizzes WHERE id = $1", id).Scan(&quiz.ID, &quiz.Title, &quiz.AuthorID, &quiz.Rating)
	if err != nil {
		return nil, fmt.Errorf("error getting quiz: %w", err)
	}

	rows, err := q.db.Query("SELECT id, text FROM questions WHERE quiz_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting questions: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		question := model.Question{}
		err := rows.Scan(&question.ID, &question.Text)
		if err != nil {
			return nil, fmt.Errorf("error scanning question: %w", err)
		}

		options, err := q.getOptionsByQuestionID(question.ID)
		if err != nil {
			return nil, fmt.Errorf("error getting options: %w", err)
		}
		question.Options = options
		quiz.Questions = append(quiz.Questions, question)
	}

	return quiz, nil	
}

func (q *QuizService) getOptionsByQuestionID(id uint) ([]model.Option, error) {
	rows, err := q.db.Query("SELECT id, text, is_correct FROM options WHERE question_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting options: %w", err)
	}
	defer rows.Close()

	var options []model.Option
	for rows.Next() {
		option := model.Option{}
		err := rows.Scan(&option.ID, &option.Text, &option.IsCorrect)
		if err != nil {
			return nil, fmt.Errorf("error scanning option: %w", err)
		}
		options = append(options, option)
	}

	return options, nil
}

