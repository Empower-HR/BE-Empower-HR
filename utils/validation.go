package utils

import (
	"errors"
	"math/rand"
	"regexp"
	"time"
)

type AccountUtilityInterface interface {
	EmailValidator(inputEmail string) error
	PasswordValidator(inputPassword string) error
	PhoneNumberValidator(inputHP string) error
	GeneratePassword(length int) (string, error)
	NumberLoop(n int) ([]int, error)
}

type accountUtility struct{}

func NewAccountUtility() AccountUtilityInterface {
	return &accountUtility{}
}

// EmailValidator implements AccountUtilityInterface.
func (ac *accountUtility) EmailValidator(inputEmail string) error {
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@gmail\.com$`, inputEmail)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("email harus menggunakan domain gmail.com")
	}
	return nil
}

// PasswordValidator implements AccountUtilityInterface.
func (ac *accountUtility) PasswordValidator(inputPassword string) error {
	if len(inputPassword) < 8 {
		return errors.New("password harus terdiri dari minimal 8 karakter")
	}
	matched, err := regexp.MatchString(`[A-Za-z]`, inputPassword)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("password harus mengandung huruf")
	}
	matched, err = regexp.MatchString(`[0-9]`, inputPassword)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("password harus mengandung angka")
	}
	return nil
}

// PhoneNumberValidator implements AccountUtilityInterface.
func (ac *accountUtility) PhoneNumberValidator(inputHP string) error {
	matched, err := regexp.MatchString(`^\d{1,12}$`, inputHP)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("nomor telepon harus terdiri dari maksimal 12 angka")
	}
	return nil
}

// GeneratePassword implements AccountUtilityInterface.
func (*accountUtility) GeneratePassword(length int) (string, error) {
	if length < 8 {
		return "", errors.New("password harus terdiri dari minimal 8 karakter")
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}

	// Ensure the password meets the validation criteria
	passwordStr := string(password)
	if err := (&accountUtility{}).PasswordValidator(passwordStr); err != nil {
		return "", err
	}

	return passwordStr, nil
}

// NumberLoop implements AccountUtilityInterface.
func (ac *accountUtility) NumberLoop(n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("input harus lebih besar dari 0")
	}

	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result, nil
}
