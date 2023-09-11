package main

import (
	"fmt"
	interpreter "github.com/noahrepublic/gosharp/intepreter"
	"os"

	"github.com/noahrepublic/gosharp/parser"
)

func main() {

	if len(os.Args) == 1 || len(os.Args) > 2 {
		panic("Invalid number of arguments")
	}

	passedFile := os.Args[1]

	data, err := os.ReadFile(passedFile)

	if err != nil {
		panic(err)
	}

	if data == nil {
		panic("Empty or invalid file")
	}

	parser := parser.Parser{}
	program := parser.Create(string(data))

	result := interpreter.EvalProgram(*program)

	fmt.Print(result, "\n")
}
