package utility

import (
	"errors"
	"regexp"
	"strings"
)

// Validate checks if the given email, name, password, and username meet the requirements.
func Validate(email, name, password, username string) error {
	if !isValidEmail(email) {
		return errors.New("invalid email address")
	}
	if !isValidName(name) {
		return errors.New("invalid name: name must be at least 3 characters, and contain only letters, spaces, and hyphens")
	}
	if !isValidPassword(password) {
		return errors.New("invalid password: password must be at least 8 characters and contain at least one uppercase, one lowercase, and one number")
	}
	if !isValidUsername(username) {
		return errors.New("invalid username: username must be at least 4 characters, and contain only letters, numbers, and underscores")
	}
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidName(name string) bool {
	nameRegex := `^[a-zA-Z\s-]{3,}$`
	re := regexp.MustCompile(nameRegex)
	return re.MatchString(name)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper, hasLower, hasDigit := false, false, false
	for _, c := range password {
		if !hasUpper && strings.ToUpper(string(c)) == string(c) && strings.ToLower(string(c)) != string(c) {
			hasUpper = true
		}
		if !hasLower && strings.ToUpper(string(c)) != string(c) && strings.ToLower(string(c)) == string(c) {
			hasLower = true
		}
		if !hasDigit && strings.ContainsAny(string(c), "0123456789") {
			hasDigit = true
		}
	}
	return hasUpper && hasLower && hasDigit
}

func isValidUsername(username string) bool {
	usernameRegex := `^[\w]{4,}$`
	re := regexp.MustCompile(usernameRegex)
	return re.MatchString(username)
}