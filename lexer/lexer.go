package lexer

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	Number     TokenType = iota
	Identifier TokenType = iota
	Equals     TokenType = iota

	Operation TokenType = iota

	OpenParenthesis  TokenType = iota
	CloseParenthesis TokenType = iota

	Let TokenType = iota

	Exit TokenType = iota // a token to signify the end for parsers
	Null TokenType = iota
)

type Token struct {
	Value string
	Type  TokenType
}

var tokenMap = map[string]TokenType{
	"=": Equals,

	"(": OpenParenthesis,
	")": CloseParenthesis,

	"+": Operation,
	"-": Operation,
	"*": Operation,
	"/": Operation,
	"%": Operation,
}

var identifierMap = map[string]TokenType{
	"let": Let,

	"null": Null,
}

func isNumber(character string) bool {
	_, err := strconv.Atoi(character)

	return err == nil
}

func isLetter(character string) bool {
	for _, rune := range character { // always ONE iteration.. surely a better way
		if !unicode.IsLetter(rune) {
			return false
		}
	}

	return true
}

func Tokenize(source string) (tokens []Token) {
	tokens = make([]Token, 0)

	skip := 0 // if you have to skip a character, this is how many you skip

	lineCount := 0
	rowCount := 0

	for i, character := range source {
		rowCount++
		if skip > 0 {
			skip--
			continue
		}

		data := string(character)

		if data == " " {
			continue
		}

		if data == "\n" {
			rowCount = 0
			lineCount++
			continue
		}

		token := tokenMap[string(character)]

		if token != 0 {
			tokens = append(tokens, Token{Value: data, Type: token})
			continue
		}

		// tokenize multichar tokens

		numberIdentifier := isNumber(data)

		if numberIdentifier {
			// loop through to get the ENTIRE number

			var number string

			var length int = 0

			for _, src := range source[i:] {
				character := string(src)

				if isNumber(character) {

					number += character
					length++
					continue
				}

				break
			}

			tokens = append(tokens, Token{Value: number, Type: Number})

			skip = length - 1
			continue
		}

		letterIdentifier := isLetter(data)

		if !letterIdentifier {
			panic(fmt.Sprintf("Invalid identifier at line: %d:%d", lineCount, rowCount))
		}

		var identifier string

		var length int

		for _, src := range source[i:] {
			character := string(src)

			if isLetter(character) {
				identifier += character
				length++
				continue
			}

			break
		}

		identifierToken := identifierMap[identifier]

		if identifierToken == 0 {
			tokens = append(tokens, Token{Value: identifier, Type: Identifier}) // we push, it could be a variable, originally was panicing
		}

		tokens = append(tokens, Token{Value: identifier, Type: identifierToken})

		skip = length
	}

	tokens = append(tokens, Token{Value: "End", Type: Exit})
	return
}
