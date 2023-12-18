package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lalathealter/ysingh-backend-assignment/commlogic"
)

const TagREPL = "YSS>>: "

func main() {

	fmt.Println("Ysingh v1.0.0")

	scanner := bufio.NewScanner(os.Stdin)
	comms := commlogic.ProduceCommandInterpreter()
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
		fmt.Println(commlogic.GetWarehousesColl())
		fmt.Println(commlogic.GetCatalog())

	}
}
