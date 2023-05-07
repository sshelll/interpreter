package part8

import (
	"fmt"
	"reflect"
)

type Interpreter struct {
	parser  *Parser
	visitor *astVisitor
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{
		parser:  NewParser(NewLexer(text)),
		visitor: &astVisitor{},
	}
}

func (it *Interpreter) Interpret() float64 {
	tree := it.parser.Parse()
	return it.visitor.visit(tree)
}

type astVisitor struct{}

func (visitor *astVisitor) visit(node astNode) (result float64) {
	switch node.(type) {
	case *astBinOp:
		return visitor.visitBinOp(node)
	case *astNum:
		return visitor.visitNum(node)
	case *astUnaryOp:
		return visitor.visitUnaryOp(node)
	default:
		panic(fmt.Sprintf("invalid type of astNode: %v", reflect.TypeOf(node)))
	}
}

func (visitor *astVisitor) visitBinOp(node astNode) (result float64) {
	bo := node.(*astBinOp)

	l := visitor.visit(bo.Left)
	r := visitor.visit(bo.Right)

	switch t := bo.Op.Type; t {
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	case MUL:
		return l * r
	case DIV:
		return l / r
	default:
		panic(fmt.Sprintf("invalid op of binOp: %v", t))
	}
}

func (visitor *astVisitor) visitNum(node astNode) (result float64) {
	num := node.(*astNum)
	return num.Value.Float()
}

func (visitor *astVisitor) visitUnaryOp(node astNode) (result float64) {
	uo := node.(*astUnaryOp)
	r := visitor.visit(uo.Right)

	switch t := uo.Op.Type; t {
	case PLUS:
		return r
	case MINUS:
		return -r
	default:
		panic(fmt.Sprintf("invalid op of unaryOp: %v", t))
	}
}
