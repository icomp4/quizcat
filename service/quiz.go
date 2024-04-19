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
	quiz.AllTimePlays = 0
	quiz.DailyPlays = 0
	quiz.WeeklyPlays = 0
	quiz.MonthlyPlays = 0
	quiz.Rating = 0
	if quiz.Picture == "" {
		quiz.Picture = "https://www.jootoor.com/wp-content/uploads/2023/09/Orange-Gradient.jpg"
	}
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
	err = tx.QueryRow("INSERT INTO quizzes (title, author_id, picture, rating, amount_of_ratings, daily_plays, weekly_plays, monthly_plays, all_time_plays) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", quiz.Title, quiz.AuthorID, quiz.Picture,0 , 0, 0, 0, 0, 0).Scan(&quizID)
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
	for _, category := range quiz.Categories {
		var categoryID uint
		err = tx.QueryRow("SELECT id FROM categories WHERE name = $1", category.Name).Scan(&categoryID)
		if err != nil {
			err = tx.QueryRow("INSERT INTO categories (name) VALUES ($1) RETURNING id", category.Name).Scan(&categoryID)
			if err != nil {
				return fmt.Errorf("error creating category: %w", err)
			}
		}

		if _, err := tx.Exec("INSERT INTO quiz_categories (quiz_id, category_id) VALUES ($1, $2)", quizID, categoryID); err != nil {
			return fmt.Errorf("error assigning category: %w", err)
		}
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
func (q *QuizService) GetQuizByID(id int) (*model.Quiz, error) {
	quiz := &model.Quiz{}
	err := q.db.QueryRow("SELECT id, title, author_id, picture, rating, daily_plays, weekly_plays, monthly_plays, all_time_plays, created_at FROM quizzes WHERE id = $1", id).
		Scan(&quiz.ID, &quiz.Title, &quiz.AuthorID, &quiz.Picture, &quiz.Rating, &quiz.DailyPlays, &quiz.WeeklyPlays, &quiz.MonthlyPlays, &quiz.AllTimePlays, &quiz.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting quiz: %w", err)
	}

	categories, err := q.GetCategoriesByQuizID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting categories: %w", err)
	}
	quiz.Categories = categories

	rows, err := q.db.Query("SELECT id, text FROM questions WHERE quiz_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting questions: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		question := model.Question{}
		question.QuizID = uint(id)
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
		option.QuestionID = id
		err := rows.Scan(&option.ID, &option.Text, &option.IsCorrect)
		if err != nil {
			return nil, fmt.Errorf("error scanning option: %w", err)
		}
		options = append(options, option)
	}

	return options, nil
}
func (q *QuizService) GetCategoriesByQuizID(id int) ([]model.Category, error) {
	categories := []model.Category{}
	rows, err := q.db.Query("SELECT c.id, c.name FROM categories c JOIN quiz_categories qc ON c.id = qc.category_id WHERE qc.quiz_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting categories: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning category: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func (q *QuizService) IncrementPlays(id int) error {
	_, err := q.db.Exec("UPDATE quizzes SET daily_plays = daily_plays + 1, weekly_plays = weekly_plays + 1, monthly_plays = monthly_plays + 1, all_time_plays = all_time_plays + 1 WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error incrementing plays: %w", err)
	}
	return nil
}

func (q *QuizService) RateQuiz(id int, rating float32) error {
	var amountOfRatings int
	var currentRating float32
	err := q.db.QueryRow("SELECT rating, amount_of_ratings FROM quizzes WHERE id = $1", id).Scan(&currentRating, &amountOfRatings)
	if err != nil {
		return fmt.Errorf("error getting rating: %w", err)
	}
	newRating := (currentRating*float32(amountOfRatings) + rating) / float32(amountOfRatings+1)
	_, err = q.db.Exec("UPDATE quizzes SET rating = $1, amount_of_ratings = $2 WHERE id = $3", newRating, amountOfRatings+1, id)
	if err != nil {
		return fmt.Errorf("error rating quiz: %w", err)
	}
	return nil
}

func (q *QuizService) GetRating(id int) (float32, error) {
	var rating float32
	err := q.db.QueryRow("SELECT rating FROM quizzes WHERE id = $1", id).Scan(&rating)
	if err != nil {
		return 0, fmt.Errorf("error getting rating: %w", err)
	}
	return rating, nil
}

func (q *QuizService) GetTopQuizzesPerPeriod(period string) ([]model.Quiz, error) {
	if period != "daily" && period != "weekly" && period != "monthly" && period != "all_time" {
		return nil, fmt.Errorf("invalid period")
	}
	period += "_plays"
	quizzes := []model.Quiz{}

	query := fmt.Sprintf("SELECT id, title, author_id, picture, rating, daily_plays, weekly_plays, monthly_plays, all_time_plays, created_at FROM quizzes ORDER BY %s DESC LIMIT 100", period)
	rows, err := q.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting top quizzes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		
		quiz := model.Quiz{}
		err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.AuthorID, &quiz.Picture, &quiz.Rating, &quiz.DailyPlays, &quiz.WeeklyPlays, &quiz.MonthlyPlays, &quiz.AllTimePlays, &quiz.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning quiz: %w", err)
		}
		categories, err := q.GetCategoriesByQuizID(int(quiz.ID))
		if err != nil {
			return nil, fmt.Errorf("error getting categories: %w", err)
		}
		quiz.Categories = categories
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (q *QuizService) GetQuizzesByCategory(category string) ([]model.Quiz, error) {
	quizzes := []model.Quiz{}
	rows, err := q.db.Query("SELECT q.id, q.title, q.author_id, q.picture, q.rating, q.daily_plays, q.weekly_plays, q.monthly_plays, q.all_time_plays, q.created_at FROM quizzes q JOIN quiz_categories qc ON q.id = qc.quiz_id JOIN categories c ON c.id = qc.category_id WHERE c.name = $1", category)
	if err != nil {
		return nil, fmt.Errorf("error getting quizzes: %w", err)
	}
	for rows.Next() {
		quiz := model.Quiz{}
		err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.AuthorID, &quiz.Picture, &quiz.Rating, &quiz.DailyPlays, &quiz.WeeklyPlays, &quiz.MonthlyPlays, &quiz.AllTimePlays, &quiz.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning quiz: %w", err)
		}
		quizzes = append(quizzes, quiz)
	}
	defer rows.Close()
	return quizzes, nil
}

func (q *QuizService) SearchQuizzes(search string) ([]model.Quiz, error) {
    quizzes := []model.Quiz{}
    rows, err := q.db.Query("SELECT id, title, author_id, picture, rating, daily_plays, weekly_plays, monthly_plays, all_time_plays, created_at FROM quizzes WHERE LOWER(title) LIKE $1" , "%" + search + "%")
    if err != nil {
        return nil, fmt.Errorf("error getting quizzes: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        quiz := model.Quiz{}
        err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.AuthorID, &quiz.Picture, &quiz.Rating, &quiz.DailyPlays, &quiz.WeeklyPlays, &quiz.MonthlyPlays, &quiz.AllTimePlays, &quiz.CreatedAt)
        if err != nil {
            return nil, fmt.Errorf("error scanning quiz: %w", err)
        }
        quizzes = append(quizzes, quiz)
    }

    return quizzes, nil
}

func (q *QuizService) DeleteQuiz(userID, id int) error {
	tx, err := q.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()
	var authorID int
	err = tx.QueryRow("SELECT author_id FROM quizzes WHERE id = $1", id).Scan(&authorID)
	if err != nil {
		return fmt.Errorf("error getting author: %w", err)
	}
	if authorID != userID {
		return fmt.Errorf("unauthorized")
	}
	_, err = tx.Exec("DELETE FROM quizzes WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting quiz: %w", err)
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil
}