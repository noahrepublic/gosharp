package interpreter

import (
	"fmt"
	"github.com/noahrepublic/gosharp/ast"
	"github.com/noahrepublic/gosharp/values"
	"math"
)

func EvalProgram(program ast.Program) interface{} {
	var statement ast.Test

	var result interface{}
	//var cast string
	for _, statement = range program.Data {

		//result, cast = evaluate(statement)
		//fmt.Print("Result", result, cast, "\n")

		fmt.Print(statement.Type, "\n")
	}

	//switch cast {
	//case "number":
	//	if num, ok := result.(values.Number); ok {
	//		result = values.Number{Value: num.Value, Type: "number"}
	//	}
	//case "null":
	//	if _, ok := result.(values.Null); ok {
	//		result = values.Null{Value: "null", Type: "null"}
	//	}
	//
	//}

	return result
}

func evaluateBinaryExpression(astNode ast.Test) any {

	fmt.Print("eval binary \n")

	data1, resultType := evaluate(*astNode.X)
	data2, _ := evaluate(*astNode.Y)

	var num1, num2 values.Number

	switch resultType {
	case "number":
		if n1, ok := data1.(values.Number); ok {
			if n2, ok := data2.(values.Number); ok {
				num1 = n1
				num2 = n2

				fmt.Print("Number", num1.Value, num2.Value, "\n")

			}
		}
	default:
		panic("Invalid type")
	}

	if num1.Type != "number" || num2.Type != "number" {
		fmt.Print("not numbers \n")
		return values.Null{Value: "null", Type: "null"}
	}

	var result float64 = 0

	fmt.Print("Test", num1, num2)

	switch astNode.Operator {
	case "+":
		result = num1.Value + num2.Value
	case "-":
		result = num1.Value - num2.Value
	case "*":
		result = num1.Value * num2.Value
	case "/":
		if num2.Value == 0 {
			panic("Division by zero")
		}

		result = num1.Value / num2.Value
	case "%":
		result = math.Mod(num1.Value, num2.Value)

	}

	fmt.Print("Result", num1.Value, num2.Value, result, "\n")
	return values.Number{Value: result, Type: "number"}
}

func evaluate(astNode ast.Test) (interface{}, string) {

	//fmt.Print("Evaluating: ", astNode.Type, "\n")

	switch astNode.Type {
	case ast.NumberLiteralNodeType:
		return values.Number{Value: astNode.NumValue, Type: "number"}, "number"

	case ast.BinaryExpressionNodeType:
		return evaluateBinaryExpression(astNode), "number"

	case ast.NullNodeType:
		return values.Null{Value: "null", Type: "null"}, "null"
	default:
		panic("Invalid AST node type")
	}
}
