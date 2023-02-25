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
	hostname, err := e.read("Hostname: ")
	if err != nil {
		return err
	}

	masterKey, err := e.readMasterKey(hostname)
	if err != nil {
		return err
	}

	secret, err := e.readSecret("Secret: ")
	if err != nil {
		return err
	}

	encodedSecret, err := security.Encode(masterKey, secret)
	if err != nil {
		return err
	}

	e.store.Store(string(hostname), encodedSecret)
	if err != nil {
		return err
	}

	fmt.Printf("Saved secret for hostname %s\n", hostname)
	return nil
}

func (e *Engine) LoadSecret() error {
	hostname, err := e.read("Hostname: ")
	if err != nil {
		return err
	}

	encodedSecret, err := e.store.Retrieve(string(hostname))
	if err != nil {
		return err
	}
	if encodedSecret == nil {
		fmt.Printf("Hostname '%s' has not yet been saved\n", hostname)
		return nil
	}

	masterKey, err := e.readMasterKey(hostname)
	if err != nil {
		return err
	}

	decodedSecret, err := security.Decode(masterKey, encodedSecret)
	if err != nil {
		return err
	}

	fmt.Printf("Secret: %s\n", decodedSecret)
	return nil
}

func (e *Engine) DeleteSecret() error {
	hostname, err := e.read("Hostname: ")
	if err != nil {
		return err
	}

	err = e.store.Delete(string(hostname))
	if err != nil {
		return err
	}

	fmt.Printf("Deleted secret for hostname %s\n", hostname)
	return nil
}

func (e *Engine) readMasterKey(hostname []byte) ([]byte, error) {
	masterKey, err := e.readSecret("Master key: ")
	if err != nil {
		return masterKey, err
	}
	return security.Hash32(masterKey, hostname), nil
}

func (e *Engine) read(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	str, err := e.terminal.ReadLine()
	if err != nil {
		return nil, err
	}
	return []byte(strings.TrimSpace(str)), nil
}

func (e *Engine) readSecret(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	secret, err := e.terminal.ReadPassword("> ")
	return []byte(secret), err
}
