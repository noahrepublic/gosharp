package parser

import (
	"fmt"
	"strconv"

	"github.com/noahrepublic/gosharp/ast"
	"github.com/noahrepublic/gosharp/lexer"
)

type Parser struct {
	tokens []lexer.Token
	at     int
}

func (p *Parser) pop() lexer.Token {
	var token = p.tokens[0]
	p.tokens = p.tokens[1:]
	p.at += 1
	return token
}

func (p *Parser) parseStatement() ast.Test {
	return p.parseExpression()
}

func (p *Parser) parseExpression() ast.Test {
	return p.parseAddition()
}

func (p *Parser) parseAddition() ast.Test {

	var left = p.parseMultiplicative() // deepens the tree & order of operations (tree works upwards)

	//fmt.Print("Left", left.NumValue, " Type:", left.Type, "\n")
	for p.at < len(p.tokens) {

		var token = p.tokens[p.at]

		if token.Type != lexer.Operation {
			break
		}

		if token.Value != "+" && token.Value != "-" {
			break
		}

		operator := token.Value
		p.at++

		var right = p.parseMultiplicative()

		//fmt.Print("Right ", right.NumValue, " Operator: ", operator, "\n")

		left = ast.Test{
			Type: ast.BinaryExpressionNodeType,
			X:    &left,
			Y:    &right,

			Operator: operator,
		}

	}

	return left
}

func (p *Parser) parseMultiplicative() ast.Test {
	var left = p.parsePrimary() // deepens the tree & order of operations (tree works upwards)

	//fmt.Print("Left", left.NumValue, " Type:", left.Type, "\n")
	for p.at < len(p.tokens) {

		var token = p.tokens[p.at]

		if token.Type != lexer.Operation {
			break
		}

		if token.Value != "*" && token.Value != "/" && token.Value != "%" {
			break
		}

		operator := token.Value
		p.at++

		var right = p.parsePrimary()

		//fmt.Print("Right ", right.NumValue, " Operator: ", operator, "\n")

		left = ast.Test{
			Type: ast.BinaryExpressionNodeType,
			X:    &left,
			Y:    &right,

			Operator: operator,
		}

	}

	return left
}

func (p *Parser) parsePrimary() ast.Test {

	var token = p.tokens[p.at]

	// map the type to a ast.NodeType

	switch token.Type {
	case lexer.Identifier:

		var identifier = ast.Test{
			Symbol: token.Value,
		}

		p.at++
		identifier.Type = ast.IdentifierNodeType

		return identifier
	case lexer.Number:

		val, err := strconv.ParseFloat(token.Value, 64) // floats yay!!!

		p.at++

		if err != nil {
			panic(err)
		}

		var identifier = ast.Test{
			NumValue: val,
		}
		identifier.Type = ast.NumberLiteralNodeType

		return identifier

	case lexer.OpenParenthesis:
		p.at++
		var expression = p.parseExpression()

		p.at++

		if p.tokens[p.at].Type != lexer.CloseParenthesis {
			panic("Expected closing parenthesis")
		}

		return expression

	case lexer.Null:
		p.at++

		// identifier := ast.NullLiteral{}
		identifier := ast.Test{}

		identifier.Type = ast.NullNodeType
		identifier.Value = "null"

		return identifier

	default:
		panic(fmt.Sprintf("Unexpected token type %s", token.Value))
	}
}

func (p *Parser) Create(src string) *ast.Program {
	var tokens = lexer.Tokenize(src)

	p.tokens = tokens
	fmt.Print(p.at, "\n")

	program := ast.Program{
		Type: ast.ProgramNodeType,
		Data: make([]ast.Test, 0),
	}

	for i := 0; i < len(tokens); i++ {
		//fmt.Print(tokens[p.at].Value, "Type:", tokens[p.at].Type, "\n")

		if tokens[p.at].Type == lexer.Exit {
			break
		}

		program.Data = append(program.Data, p.parseStatement())
	}

	return &program
}
