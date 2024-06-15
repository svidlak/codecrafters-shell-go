package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
)

func main() {
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic("bad input")
		}

		getInput(input)
	}

}

func getInput(input string) {
	values := strings.Split(input, "\n")
	values = strings.Split(values[0], " ")

	command := values[0]

	val, ok := commands.Commands_list[command]

	if !ok {
		fmt.Println(command + ": command not found")
		return
	}

	val(values[1:])
}
