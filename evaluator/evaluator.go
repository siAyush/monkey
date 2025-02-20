package evaluator

import (
	"github.com/siAyush/monkey/ast"
	"github.com/siAyush/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// statementss
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
