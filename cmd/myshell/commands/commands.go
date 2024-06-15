package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandProcessor struct {
	Commands map[string]func([]string)
}

func NewCommandProcessor() *CommandProcessor {
	cp := &CommandProcessor{
		Commands: make(map[string]func([]string)),
	}
	cp.initCommands()
	return cp
}

func (cp *CommandProcessor) initCommands() {
	cp.Commands["exit"] = cp.exitFunc
	cp.Commands["echo"] = cp.echoFunc
	cp.Commands["type"] = cp.typeFunc
}

func (cp *CommandProcessor) RunCommand(command string, args []string) {
	if fn, ok := cp.Commands[command]; ok {
		fn(args)
	} else {
		fmt.Fprintf(os.Stderr, "Command '%s' not found\n", command)
	}
}

func (cp *CommandProcessor) exitFunc(input []string) {
	if len(input) < 1 {
		fmt.Fprintln(os.Stderr, "exit: missing status code")
		os.Exit(1)
	}

	statusCode, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit: invalid status code '%s'\n", input[0])
		os.Exit(1)
	}

	os.Exit(statusCode)
}

func (cp *CommandProcessor) echoFunc(input []string) {
	if len(input) < 1 {
		fmt.Fprintln(os.Stderr, "echo: nothing to echo")
		return
	}
	line := strings.Join(input, " ")
	fmt.Println(line)
}

func (cp *CommandProcessor) typeFunc(input []string) {
	if len(input) < 1 {
		fmt.Fprintln(os.Stderr, "type: missing command name")
		return
	}

	_, ok := cp.Commands[input[0]]
	if !ok {
		fmt.Printf("%s: not found\n", input[0])
		return
	}
	fmt.Printf("%s is a shell builtin\n", input[0])
}
