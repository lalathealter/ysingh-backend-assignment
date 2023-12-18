package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lalathealter/ysingh-backend-assignment/commlogic"
	"github.com/lalathealter/ysingh-backend-assignment/logstream"
)

const TagREPL = "YSS>>: "

func main() {

	fmt.Println("Ysingh v1.0.0")

	scanner := bufio.NewScanner(os.Stdin)
	comms := commlogic.ProduceCommandInterpreter()
	logStreamer := logstream.ProduceLogStreamer()
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
		logStreamer.Send(input, res)
	}
}
