package util

import (
	"errors"
	"quizcat/model"
)

func ValidateQuiz(quiz *model.Quiz) error {
	if quiz.Title == "" || len(quiz.Title) < 5 {
		return errors.New("title is too short or empty")
	}
	if quiz.Questions == nil || len(quiz.Questions) == 0 {
		return errors.New("questions are required")
	}
	for _, q := range quiz.Questions {
		if q.Text == "" || len(q.Text) < 5 {
			return errors.New("question is required")
		}
		if q.Options == nil || len(q.Options) < 2 {
			return errors.New("at least 2 options are required")
		}
		correctCount := 0
		for _, o := range q.Options {
			if o.Text == "" || len(o.Text) < 1 {
				return errors.New("option text is required")
			}
			if o.IsCorrect {
				correctCount++
			}
		}
		if correctCount != 1 {
			return errors.New("exactly 1 correct option is required")
		}
	}
	return nil
}

func ValidateUser(model *model.User) error{
	if model.Username == "" || len(model.Username) < 5 {
		return errors.New("name is too short or empty")
	}
	if model.Email == "" || len(model.Email) < 5 {
		return errors.New("email is too short or empty")
	}
	if model.Password == "" || len(model.Password) < 5 {
		return errors.New("password is too short or empty")
	}
	return nil
}
