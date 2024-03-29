package service

import (
	"database/sql"
	"fmt"
	"quizcat/model"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db}
}

func (u *UserService) CreateUser(user *model.User) error {
	if user.Email == "" || user.Username == "" || user.Password == "" {
		return fmt.Errorf("missing required fields")
	}
	if err := u.db.QueryRow("INSERT INTO users (email, username, password) VALUES ($1, $2, $3)", user.Email, user.Username, user.Password).Err(); err != nil{
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
