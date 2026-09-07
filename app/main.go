package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
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
		cmd, args := tokens[0], tokens[1:]

		if cmd == "type" {
			if slices.Contains(builtin, tokens[1]) {
				fmt.Println(tokens[1] + " is a shell builtin")
			} else if path, err := exec.LookPath(args[0]); err == nil { // Implement PATH traversal logic
				fmt.Println(args[0] + " is " + path)
			} else if cmd == "type" {
				fmt.Println(tokens[1] + ": not found")
			}  else {
				fmt.Println(command[:len(command)-1] + ": command not found")
			}
		} else if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo") {
			fmt.Println(command[5:])
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

	}
	
}
