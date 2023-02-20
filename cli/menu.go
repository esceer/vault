package cli

import (
	"bufio"
	"fmt"
	"os"
)

type Menu struct {
	Options map[string]func(*bufio.Scanner)
}

func NewMenu() *Menu {
	return &Menu{Options: map[string]func(s *bufio.Scanner){
		"save":   saveSecret,
		"load":   loadSecret,
		"delete": deleteSecret,
	}}
}

func (m *Menu) Display() {
	fmt.Println("Select from the below options:")
	for key, _ := range m.Options {
		fmt.Printf("'%v'\n", key)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		selectedFunction := m.Options[scanner.Text()]
		if selectedFunction != nil {
			selectedFunction(scanner)
		} else {
			fmt.Println("The option you selected is invalid. Please try again")
		}
	}
}

func saveSecret(scanner *bufio.Scanner) {
	fmt.Println("Saving secret...")

	readInput(scanner, "Host: ")
	host := scanner.Text()

	readInput(scanner, "Secret: ")
	secret:= scanner.Text()

	Saving...
}

func loadSecret(scanner *bufio.Scanner) {}

func deleteSecret(scanner *bufio.Scanner) {}

func readInput(scanner *bufio.Scanner, displayText string) {
	fmt.Printf("%v", displayText)
	scanner.Scan()
	fmt.Println()
}
