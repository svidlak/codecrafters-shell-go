package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Commands_list = map[string]func([]string){
	"exit": exit,
	"echo": echo,
}

func exit(input []string) {
	s := input[0]

	statusCode, err := strconv.Atoi(s)
	if err != nil {
		panic("exit panic")
	}
	os.Exit(statusCode)
}

func echo(input []string) {
	line := strings.Join(input, " ")
	fmt.Print(line + "\n")
}
