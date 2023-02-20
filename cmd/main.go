package main

import (
	"fmt"

	"github.com/esceer/vault/security"
	"github.com/esceer/vault/storage"
)

func main() {
	store := storage.New()

	key := getKey()
	password := requestPassword()
	fmt.Printf("Password: %x (%v)\n", password, string(password))

	encodedPassword, _ := security.Encode(key, password)
	fmt.Printf("Encoded: %x (%v)\n", encodedPassword, string(encodedPassword))

	decodedPassword, _ := security.Decode(key, encodedPassword)
	fmt.Printf("Decoded: %x (%v)\n", decodedPassword, string(decodedPassword))
}

func getKey() []byte {
	return []byte("supersecret32byteslongcipherkey!")
}

func requestPassword() []byte {
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Enter the password...")
	// scanner.Scan()
	// return scanner.Text()
	return []byte("alma%12!0")
}
