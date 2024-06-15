package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	replMessage()
	for scanner.Scan() {
		command := cleanInput(scanner.Text())

		if res, exists := commands()[command]; exists {
			err := res.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", command)
		}
		fmt.Println()
		replMessage()
	}
}

func cleanInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func replMessage() {
	fmt.Print("Pokedex > ")
}
