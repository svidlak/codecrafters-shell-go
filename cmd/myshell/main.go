package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic("bad input")
	}
	getInput(input)
}

func getInput(input string) {
	val, ok := commands.Commands_list[input]

	if !ok {
		fmt.Println(input + ": command not found")
		return
	}

	fmt.Println(val)
}
