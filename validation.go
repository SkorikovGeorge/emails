package main

import (
	"errors"
	"strings"
)

const ErrInvalidEmail = "invalid email"

func IsValidEmail(email string) error {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 || atIndex != strings.LastIndex(email, "@") {
		return errors.New(ErrInvalidEmail)
	}

	local := email[:atIndex]
	domain := email[atIndex+1:]

	if len(local) == 0 || len(domain) == 0 {
		return errors.New(ErrInvalidEmail)
	}

	if !strings.Contains(domain, ".") {
		return errors.New(ErrInvalidEmail)
	}

	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return errors.New(ErrInvalidEmail)
	}

	return nil
}
