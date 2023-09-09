package lexer

import "fmt"

type TokenType int

const (
	Number     TokenType = iota
	Identifier TokenType = iota
	Equals     TokenType = iota

	OpenParenthesis  TokenType = iota
	CloseParenthesis TokenType = iota

	Let TokenType = iota
)

type Token struct {
	value string
	token TokenType
}

var tokenMap = map[string]TokenType{
	"let": Let,
	"=":   Equals,

	"(": OpenParenthesis,
	")": CloseParenthesis,
}

func tokenize(source string) (tokens []Token) {
	tokens = make([]Token, 0)

	for i := 0; i < len(source); i++ {
		if source[i] == ' ' {
			continue
		}

		token := tokenMap[string(source[i])]

		fmt.Println(token)

	}

	return
}
