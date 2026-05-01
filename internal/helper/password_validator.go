package helper

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	var err error
	if len(password) < 12 {
		err = errors.New("password must be at least 12 characters")
		return err
	}

	// minimal 1 huruf kapital
	upper := regexp.MustCompile(`[A-Z]`)
	if !upper.MatchString(password) {
		err = errors.New("password must contain at least 1 uppercase letter")
		return err
	}

	// minimal 1 angka
	number := regexp.MustCompile(`[0-9]`)
	if !number.MatchString(password) {
		err = errors.New("password must contain at least 1 number")
		return err
	}

	// minimal 1 simbol
	symbol := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",.<>/?\\|]`)
	if !symbol.MatchString(password) {
		err = errors.New("password must contain at least 1 special character")
		return err
	}

	return nil
}
