package service

import (
	"database/sql"
	"fmt"
	"quizcat/model"

	"golang.org/x/crypto/bcrypt"
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
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	user.Password = string(password)
	if err := u.db.QueryRow("INSERT INTO users (email, username, password) VALUES ($1, $2, $3)", user.Email, user.Username, user.Password).Err(); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (u *UserService) Login(username, password string) (int, error) {
	var hashedPassword string
	var id int
	if err := u.db.QueryRow("SELECT id, password FROM users WHERE username = $1", username).Scan(&id, &hashedPassword); err != nil {
		return -1, fmt.Errorf("error getting user: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return -1, fmt.Errorf("error comparing password: %w", err)
	}
	return id, nil
}
