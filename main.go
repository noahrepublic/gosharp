package main

import (
	"os"
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

	print(string(data))

	if data == nil {
		panic("Empty or invalid file")
	}

	tokens := lexer.tokenize(string(data))

	for _, token := range tokens {
		print(token.value)
	}
}