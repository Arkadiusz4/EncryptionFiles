# Encryption Files
 
## Overview
This is a simple command-line file encryption tool written in Golang. It allows you to encrypt and decrypt files using a password. 

## Usage
To use the tool, run the following commands:

### Encrypt a file
```go
go run . encrypt /path/to/your/file
```

This command encrypts the specified file using a password that you will be prompted to enter.

### Decrypt a file
```go
go run . decrypt /path/to/your/file
```

This command attempts to decrypt the specified file using a password that you will be prompted to enter.

### Help
```go
go run . help
```
Displays help text with usage instructions.

## Commands
- **encrypt**: Encrypts a file given a password.
- **decrypt**: Tries to decrypt a file using a password.
- **help**: Displays help text.

## Note
Make sure to keep your passwords secure, and remember that if you forget the password used for encryption, you won't be able to decrypt the file.

Feel free to explore and customize the tool according to your needs. 