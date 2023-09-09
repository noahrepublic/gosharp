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
)

type Token struct {
	Value string
	Token TokenType
}

var tokenMap = map[string]TokenType{
	"=": Equals,

	"(": OpenParenthesis,
	")": CloseParenthesis,

	"+": Operation,
	"-": Operation,
	"*": Operation,
	"/": Operation,
}

var identifierMap = map[string]TokenType{
	"let": Let,
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
			tokens = append(tokens, Token{Value: data, Token: token})
			continue
		}

		// tokenize multichar tokens

		numberIdentifier := isNumber(data)

		if numberIdentifier {
			// loop through to get the ENTIRE number

			var number string

			var length int

			for _, src := range source[i:] {
				character := string(src)

				if isNumber(character) {
					number += character
					length++
					continue
				}

				break
			}

			tokens = append(tokens, Token{Value: number, Token: Number})

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
			panic(fmt.Sprintf("Invalid identifier at line: %d:[%d:%d]", lineCount, rowCount, i+length+1)) // future: should we instead push identifier tokentype?
		}

		tokens = append(tokens, Token{Value: identifier, Token: identifierToken})

		skip = length
	}

	return
}
