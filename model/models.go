package model

type Quiz struct {
	ID       uint   `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	AuthorID uint   `db:"author_id" json:"author_id"`
	Rating   float32 `db:"rating" json:"rating"`
	Questions []Question 
}

type Question struct {
	ID     uint   `db:"id" json:"id"`
	QuizID uint   `db:"quiz_id" json:"quiz_id"`
	Text   string `db:"text" json:"text"`
	Options []Option 
}

type Option struct {
	ID       uint   `db:"id" json:"id"`
	QuestionID uint `db:"question_id" json:"question_id"`
	Text     string `db:"text" json:"text"`
	IsCorrect bool   `db:"is_correct" json:"is_correct"`
}

type User struct {
	ID       uint   `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Quizzes []Quiz 
}
