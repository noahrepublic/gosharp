package ast

// node types to build ast https://esprima.org/demo/parse.html

type NodeType string

const (
	ProgramNodeType          NodeType = "Program"
	StatementNodeType        NodeType = "Statement"
	ExpressionNodeType       NodeType = "Expression"
	BinaryExpressionNodeType NodeType = "BinaryExpression"
	IdentifierNodeType       NodeType = "Identifier"
	NumberLiteralNodeType    NodeType = "NumberLiteral"

	NullNodeType NodeType = "Null"
)

// let myVar = 1
// myVar = 1
// Implement statements

type Statement struct {
	Type NodeType
}

type ConstantRepeat struct {
	Type     NodeType
	X        *ConstantRepeat
	Y        *ConstantRepeat
	Operator string
}

type Test struct {
	Type NodeType
	Data []Test

	X        *Test
	Y        *Test
	Operator string

	Symbol string
	Value  string

	NumValue float64
}

type Program struct {
	Type NodeType
	Data []Test
}

// Expression

type Expression struct {
	Statement
}

// a + b

type BinaryExpression struct {
	Type     NodeType
	X        any // cant figure out types
	Y        any
	Operator string
}

type Identifier struct {
	Expression // Identifier
	Symbol     string
}

type NumberLiteral struct {
	Expression
	Value float64
}

type NullLiteral struct {
	Expression
	Value string
}
