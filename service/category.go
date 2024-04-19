package service

import (
	"database/sql"
	"fmt"
	"quizcat/model"
)

type CategoryService struct {
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{
		db: db}
}

func (q *CategoryService) CreateCategory(name string) error {
	_, err := q.db.Exec("INSERT INTO categories (name) VALUES ($1)", name)
	if err != nil {
		return err
	}
	return nil
}

func (q *CategoryService) GetCategories() ([]string, error) {
	rows, err := q.db.Query("SELECT name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []string
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (q *CategoryService) AssignCategoryToQuiz(userID int, quizID int, categoryID int) error {
	var quizownerID int
	err := q.db.QueryRow("SELECT author_id FROM quizzes WHERE id = $1", quizID).Scan(&quizownerID)
	if err != nil {
		return err
	}
	if quizownerID != userID {
		return fmt.Errorf("only the author of the quiz can assign categories to it")
	}
	_, err = q.db.Exec("INSERT INTO quiz_categories (quiz_id, category_id) VALUES ($1, $2)", quizID, categoryID)
	if err != nil {
		return err
	}
	return nil
}

func (q *CategoryService) GetQuizzesByCategory(category string) ([]model.Quiz, error) {
	var quizzes []model.Quiz

	query := `
        SELECT q.id, q.title, q.author_id, q.picture, q.rating, q.amount_of_ratings, q.daily_plays, 
               q.weekly_plays, q.monthly_plays, q.all_time_plays
        FROM quizzes q
        JOIN quiz_categories qc ON q.id = qc.quiz_id
        JOIN categories c ON qc.category_id = c.id
        WHERE LOWER(c.name) = $1
    `

	rows, err := q.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var quiz model.Quiz
		err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.AuthorID, &quiz.Picture, &quiz.Rating, &quiz.AmountOfRatings,
			&quiz.DailyPlays, &quiz.WeeklyPlays, &quiz.MonthlyPlays, &quiz.AllTimePlays)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}
