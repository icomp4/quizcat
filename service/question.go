package service

import "database/sql"

type QuestionService struct {
	db *sql.DB
}

func NewQuestionService(db *sql.DB) *QuestionService {
	return &QuestionService{
		db: db}
}

func (q *QuestionService) CreateQuestion(userid, quizID int, questionText string) error {
	_, err := q.db.Exec(`
		INSERT INTO questions (quiz_id, text) 
		VALUES ($1, $2)`, 
		quizID, questionText)
	if err != nil {
		return err
	}
	return nil

}
func(q *QuestionService) EditQuestionText(userid int, questionID int, questionText string) error {
    _, err := q.db.Exec(`
        UPDATE questions 
        SET text = $1 
        FROM quizzes 
        WHERE questions.id = $2 
        AND quizzes.author_id = $3 
        AND quizzes.id = questions.quiz_id`, 
        questionText, questionID, userid)
    if err != nil {
        return err
    }
    return nil
}

func(q *QuestionService) DeleteQuestion(userid int, questionID int) error {
	_, err := q.db.Exec(`
		DELETE FROM questions 
		USING quizzes 
		WHERE questions.id = $1 
		AND quizzes.author_id = $2 
		AND quizzes.id = questions.quiz_id`, 
		questionID, userid)
	if err != nil {
		return err
	}
	return nil
}

func(q *QuestionService) EditOptionText(userid int, optionID int, optionText string) error {
    _, err := q.db.Exec(`
        UPDATE options 
        SET text = $1 
        WHERE id = $2 AND question_id IN (
            SELECT questions.id 
            FROM questions 
            JOIN quizzes ON quizzes.id = questions.quiz_id
            WHERE quizzes.author_id = $3
        )`, 
        optionText, optionID, userid)
    if err != nil {
        return err
    }
    return nil
}

func(q *QuestionService) EditOptionCorrect(userid int, optionID int, correct bool) error {
	_, err := q.db.Exec(`
		UPDATE options 
		SET is_correct = $1 
		WHERE id = $2 AND question_id IN (
			SELECT questions.id 
			FROM questions 
			JOIN quizzes ON quizzes.id = questions.quiz_id
			WHERE quizzes.author_id = $3
		)`, 
		correct, optionID, userid)
	if err != nil {
		return err
	}
	return nil
}

func(q *QuestionService) DeleteOption(userid, optionID int) error {
    _, err := q.db.Exec(`
        DELETE FROM options 
        USING questions, quizzes 
        WHERE options.id = $1 
        AND questions.id = options.question_id
        AND quizzes.id = questions.quiz_id
        AND quizzes.author_id = $2`, 
        optionID, userid)
    if err != nil {
        return err
    }
    return nil
}
func(q *QuestionService) CreateOption(userid, questionID int, optionText string, isCorrect bool) error {
	_, err := q.db.Exec(`
		INSERT INTO options (question_id, text, is_correct) 
		VALUES ($1, $2, $3)`, 
		questionID, optionText, isCorrect)
	if err != nil {
		return err
	}
	return nil
}