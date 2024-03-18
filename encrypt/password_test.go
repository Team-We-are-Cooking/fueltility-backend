package encrypt

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_PasswordHashing(t *testing.T) {
	data := []struct {
		name     string
		password string
	}{
		{"empty", ""},
		{"lowercase", "ad2323"},
		{"space", "ab  c123"},
		{"uppercase", "AB123D"},
		{"mix", "@Pw__123"},
		{"mix", "#Password@123"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			hashedPassword, err := HashPassword(d.password)

			if err != nil {
				t.Fatalf("%s was unable to be hashed: %s", d.password, err.Error())
			}

			if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(d.password)); err != nil {
				t.Errorf("%s was incorrectly hashed", d.password)
			}
		})
	}
}
