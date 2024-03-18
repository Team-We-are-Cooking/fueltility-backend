package encrypt

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	if utf8.RuneCountInString(password) == 0 {
		return nil, fmt.Errorf("password cannot be an empty string")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("unable to hash password, %w", err)
	}

	return hash, nil
}
