package main

import (
	"sync"

	"github.com/esceer/vault/cli"
	"github.com/esceer/vault/storage"
)

var (
	once   sync.Once
	engine *cli.Engine
)

func GetEngine() *cli.Engine {
	once.Do(func() {
		store := storage.New()
		engine = cli.NewEngine(store)
	})
	return engine
}

func main() {
	cli.DisplayMenu(GetEngine())

	// key := getKey()
	// password := requestPassword()
	// fmt.Printf("Password: %x (%v)\n", password, string(password))

	// encodedPassword, _ := security.Encode(key, password)
	// fmt.Printf("Encoded: %x (%v)\n", encodedPassword, string(encodedPassword))

	// decodedPassword, _ := security.Decode(key, encodedPassword)
	// fmt.Printf("Decoded: %x (%v)\n", decodedPassword, string(decodedPassword))
}

// func getKey() []byte {
// 	return []byte("supersecret32byteslongcipherkey!")
// }

// func requestPassword() []byte {
// 	return []byte("alma%12!0")
// }
