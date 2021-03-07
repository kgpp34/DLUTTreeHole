package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)


var (
	rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")
	rxUsername = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{0,17}$")
	// ErrUserNotFound used when the user wasn't found on the db
	ErrUserNotFound = errors.New("user not found")
	// ErrInvalidEmail stands for invalid email when user begin to login 
	ErrInvalidEmail = errors.New("invalid email")
	// ErrInvalidUsername used when the username is not valid
	ErrInvalidUsername = errors.New("invalid username")
	// ErrEmailTaken used when there is already an user registerer with that email
	ErrEmailTaken = errors.New("email taken")
	// ErrUsernameTaken used when there is already an user registered with that username
	ErrUsernameTaken = errors.New("username taken")
)

// User model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// CreateUser inserts a user in the database
func (s *Service) CreateUser(ctx context.Context, email, username string) error {
	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return ErrInvalidEmail
	}
	
	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username) {
		return ErrInvalidUsername
	}

	query := "insert into users (email, username) values($1, $2)"
	_, err := s.db.ExecContext(ctx, query, email, username)
	unique := isUniqueViolation(err)

	if unique && strings.Contains(err.Error(), "email") {
		return ErrEmailTaken
	}
	
	if unique && strings.Contains(err.Error(), "username") {
		return ErrUsernameTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}
	
	return nil
}