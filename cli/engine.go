package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/esceer/vault/security"
	"github.com/esceer/vault/storage"
	"golang.org/x/term"
)

type Engine struct {
	store    storage.IStore
	terminal *term.Terminal
}

func NewEngine(store storage.IStore) *Engine {
	return &Engine{
		store:    store,
		terminal: term.NewTerminal(os.Stdin, "> "),
	}
}

func (e *Engine) SaveSecret() error {
	fmt.Println("Saving secret...")

	fmt.Print("Master key: ")
	masterKey, err := e.readSecret()
	if err != nil {
		fmt.Println("Failed reading master key")
		return err
	}
	fmt.Println()

	hostname, err := e.read()
	if err != nil {
		fmt.Println("Failed reading hostname")
		return err
	}

	secret, err := e.readSecret()
	if err != nil {
		fmt.Println("Failed reading secret")
		return err
	}

	encodedSecret, err := security.Encode(masterKey, secret)
	if err != nil {
		fmt.Println("Failed encoding secret")
		return err
	}

	e.store.Store(string(hostname), encodedSecret)

	fmt.Printf("Saved secret for hostname %s", hostname)
	return nil
}

func (e *Engine) LoadSecret() error {
	fmt.Println("Loading secret...")

	fmt.Print("Master key: ")
	masterKey, err := e.readSecret()
	if err != nil {
		fmt.Println("Failed reading master key")
		return err
	}
	fmt.Println()

	hostname, err := e.read()
	if err != nil {
		fmt.Println("Failed reading hostname")
		return err
	}

	encodedSecret, err := e.store.Retrieve(string(hostname))
	if err != nil {
		fmt.Println("Failed retrieving secret")
		return err
	}

	decodedSecret, err := security.Decode(masterKey, encodedSecret)
	if err != nil {
		fmt.Println("Failed decoding secret")
		return err
	}

	fmt.Printf("Secret: %s", decodedSecret)
	return nil
}

func (e *Engine) DeleteSecret() error {
	fmt.Println("Deleting secret...")

	hostname, err := e.read()
	if err != nil {
		fmt.Println("Failed reading hostname")
		return err
	}

	err = e.store.Delete(string(hostname))
	if err != nil {
		fmt.Println("Failed deleting secret")
		return err
	}

	fmt.Printf("Deleted secret for hostname %s", hostname)
	return nil
}

func (e *Engine) read() ([]byte, error) {
	str, err := e.terminal.ReadLine()
	if err != nil {
		fmt.Println("Failed to read from terminal")
		return nil, err
	}
	return []byte(strings.TrimSpace(str)), nil
}

func (e *Engine) readSecret() ([]byte, error) {
	secret, err := e.terminal.ReadPassword("> ")
	return []byte(secret), err
}
