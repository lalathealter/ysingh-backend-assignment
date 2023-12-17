package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const TagREPL = "YSS>>: "

func main() {

	fmt.Println("Ysingh v1.0.0")

	scanner := bufio.NewScanner(os.Stdin)
	comms := ProduceCommandInterpreter()
	for {
		fmt.Print(TagREPL)
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		args := strings.Split(input, " ")
		res, err := comms.Execute(args)
		if err != nil {
			res = err.Error()
		}
		fmt.Println(res)

	}
}

type CommandInterpreter map[string]func(any) string

func (ci CommandInterpreter) Execute(args []string) (string, error) {
	comfun, ok := ci[args[0]]
	if !ok {
		return "", errors.New("UKNOWN COMMAND")
	}
	ans := comfun(args)
	return ans, nil
}

func ProduceCommandInterpreter() CommandInterpreter {
	return CommandInterpreter{
		"STOP": ExitApp,
	}
}

func ExitApp(_ any) string {
	go os.Exit(0)
	return io.EOF.Error()
}
