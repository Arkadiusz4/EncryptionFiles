package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file, and decrypt to decrypt a file")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("File encryption")
	fmt.Println("Simple file encrypter for your day-to-day needs")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\t go run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\t encrypt \t Encrypts a file given a password")
	fmt.Println("\t decrypt \t Tries to decrypt a file using a password")
	fmt.Println("\t help \t\t Displays help text")
	fmt.Println("")
}

func encryptHandle() {}

func decryptHandle() {

}

func getPassword()      {}
func validatePassword() {}
func validateFile()     {}
