package cli

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayMenu(engine *Engine) {
	printAvailableOptions()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var err error
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
			fmt.Printf("Unknown option '%v'", scanner.Text())
			printAvailableOptions()
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func printAvailableOptions() {
	fmt.Println("Select from the following options: %v", []string{"save", "load", "delete", "quit"})
}
