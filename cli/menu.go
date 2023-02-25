package cli

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayMenu(engine *Engine) {
	printAvailableOptions()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var err error

		fmt.Print("> ")
		scanner.Scan()
		switch scanner.Text() {
		case "save":
			err = engine.SaveSecret()
		case "load":
			err = engine.LoadSecret()
		case "delete":
			err = engine.DeleteSecret()
		case "quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Printf("Unknown option '%v'. ", scanner.Text())
			printAvailableOptions()
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}

func printAvailableOptions() {
	fmt.Println("Select from the following options:")
	for _, option := range []string{"save", "load", "delete", "quit"} {
		fmt.Printf("- %v\n", option)
	}
	fmt.Println()
}
