package main

import (
	"fmt"
	"os"

	"github.com/noahrepublic/gosharp/lexer"
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

	tokens := lexer.Tokenize(string(data))

	for _, token := range tokens {
		fmt.Println("Token", token.Value)
	}
}
