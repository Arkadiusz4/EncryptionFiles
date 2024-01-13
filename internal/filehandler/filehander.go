package filehandler

import (
	"fmt"
	"github.com/Arkadiusz4/EncryptionFiles/internal/filecrypt"
	"golang.org/x/term"
	"os"
)

func HandleEncrypt() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("File successfully protected.")
}

func HandleDecrypt() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter a password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("File successfully decrypted.")
}
