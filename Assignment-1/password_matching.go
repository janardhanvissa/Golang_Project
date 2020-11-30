package main

import (
	"errors"
	"fmt"
	"unicode"
)

var ErrInvalidPassword = errors.New(`The password should at least have 7 letters, at least 1 number, at least 1 upper case, at least 1 special character.`)

func VerifyPassword(pw string) error {
	if len(pw) < 10 {
		return ErrInvalidPassword
	}
	var num, lower, upper, spec bool
	for _, r := range pw {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			spec = true
		}
	}
	if num && lower && upper && spec {
		return nil
	}
	return ErrInvalidPassword
}
func main() {
	fmt.Println(VerifyPassword("Janardhan@19"))
	fmt.Println(VerifyPassword("Janardhan@19"))
	fmt.Println("Both the passwords are matching")
}
