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

func (u *UserService) GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow("SELECT id, email, username FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Username); err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	row, err := u.db.Query("SELECT id FROM quizzes WHERE author_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting quizzes: %w", err)
	}
	for row.Next() {
		var quizID int
		if err := row.Scan(&quizID); err != nil {
			return nil, fmt.Errorf("error scanning quiz: %w", err)
		}
		user.Quizzes = append(user.Quizzes, model.Quiz{ID: uint(quizID)})
	}

	return &user, nil
}