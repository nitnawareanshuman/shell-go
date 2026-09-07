package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"slices"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	builtin := []string{"echo", "exit", "type"}

	for {

		fmt.Print("$ ")

		command, err :=reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		tokens := strings.Split(command, " ")

		if tokens[0] == "type" && slices.Contains(builtin, tokens[1]) {
			fmt.Println(tokens[1] + "is a shell builtin")
		} else if tokens[0] == "type" {
			fmt.Println(tokens[1] + ": not found")
		} else if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo") {
			fmt.Println(command[5:])
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

	}
	
}
