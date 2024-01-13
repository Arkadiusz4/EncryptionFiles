package filehandler

import (
	"bytes"
	"fmt"
	"golang.org/x/term"
	"os"
)

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}

	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
