package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {

	fmt.Print("$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Printf(command[:len(command)-1] + ": command not found")
}
