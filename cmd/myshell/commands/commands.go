package commands

import (
	"os"
	"strconv"
)

var Commands_list = map[string]func([]string){
	"exit": exit,
}

func exit(input []string) {
	s := input[0]

	statusCode, err := strconv.Atoi(s)
	if err != nil {
		panic("exit panic")
	}
	os.Exit(statusCode)
}
