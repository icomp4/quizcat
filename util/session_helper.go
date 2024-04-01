package util

import (
	"errors"

	"github.com/gofiber/fiber/v3/middleware/session"
)

func GetUserIDFromSession(sess session.Session) (int, error) {
	userID := sess.Get("userID")
	if userID == nil {
		return -1, errors.New("Unauthorized")
	}
	userIDInt, ok := userID.(int)
	if !ok {
		return -1, errors.New("error getting user ID")
	}
	return userIDInt, nil
}
